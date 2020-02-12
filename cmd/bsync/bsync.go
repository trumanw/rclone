package bsync

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/trumanw/rclone/cmd"
	"github.com/trumanw/rclone/fs"
	"github.com/trumanw/rclone/fs/config/flags"
	"github.com/trumanw/rclone/fs/filter"
	"github.com/trumanw/rclone/fs/fserrors"
	"github.com/trumanw/rclone/fs/operations"
	"github.com/trumanw/rclone/fs/sync"
)

var (
	createEmptySrcDirs = false
	ignoreFile         string
	syncOnlyFn         string
)

func init() {
	cmd.Root.AddCommand(commandDefinition)
	cmdFlags := commandDefinition.Flags()
	flags.BoolVarP(cmdFlags, &createEmptySrcDirs, "create-empty-src-dirs", "", createEmptySrcDirs, "Create empty source dirs on destination after sync")
	flags.StringVarP(cmdFlags, &ignoreFile, "ignore-file", "", "", "Specify the ignore file path as the filtering rules.")
	flags.StringVarP(cmdFlags, &syncOnlyFn, "sync-only-fn", "", "", "Specify the sync-only file name.")
}

var commandDefinition = &cobra.Command{
	Use:   "bsync source:path dest:path",
	Short: `Make the list of sources and dests identical, batch modifying destination only.`,
	Long: `
	Sync the source to the destination, changing the destination
	only.  Doesn't transfer unchanged files, testing by size and
	modification time or MD5SUM.  Destination is updated to match
	source, including deleting files if necessary.

	bsync s3:src/0@dst/0 s3:src/1@dst/1
	
	**Important**: Since this can cause data loss, test first with the
	` + "`" + `--dry-run` + "`" + ` flag to see exactly what would be copied and deleted.
	
	Note that files in the destination won't be deleted if there were any
	errors at any point.
	
	It is always the contents of the directory that is synced, not the
	directory so when source:path is a directory, it's the contents of
	source:path that are copied, not the directory name and contents.  See
	extended explanation in the ` + "`" + `copy` + "`" + ` command above if unsure.
	
	If dest:path doesn't exist, it is created and the source:path contents
	go there.
	
	**Note**: Use the ` + "`-P`" + `/` + "`--progress`" + ` flag to view real-time transfer statistic
`,
	Run: func(command *cobra.Command, argss []string) {
		cmd.CheckArgs(1, 10, command, argss)
		c := context.WithValue(context.Background(), "IgnoreFile", ignoreFile)
		c = context.WithValue(c, "SyncOnlyFn", syncOnlyFn)
		// arrays for BatchSync
		fsrcDirs := []fs.Fs{}
		fdstDirs := []fs.Fs{}
		// arrays for BatchCopyFiles
		fsrcFileBases := []fs.Fs{}
		fdstFileBases := []fs.Fs{}
		fsrcFileNames := []string{}
		fdstFileNames := []string{}
		for _, args := range argss {
			fargs := strings.Split(args, "@")
			// fsrc, _, fdst := cmd.NewFsSrcFileDst(fargs)
			// clear filters
			filter.Active.Clear()
			filter.Active, _ = filter.NewFilter(nil)
			fsrcBaseName, fsrcFileName, fdstBaseName, fdstFileName := cmd.NewFsSrcDstFiles(fargs)
			if fsrcFileName == "" {
				fsrcDirs = append(fsrcDirs, fsrcBaseName)
				fdstDirs = append(fdstDirs, fdstBaseName)
			} else {
				fsrcFileBases = append(fsrcFileBases, fsrcBaseName)
				fdstFileBases = append(fdstFileBases, fdstBaseName)
				fsrcFileNames = append(fsrcFileNames, fsrcFileName)
				fdstFileNames = append(fdstFileNames, fdstFileName)
			}
		}
		cmd.Run(true, true, command, func() error {
			var dirSyncErr, fileSyncErr error
			if len(fsrcDirs) > 0 && len(fsrcDirs) == len(fdstDirs) {
				errfs, dirSyncErr := sync.BatchSync(c, fdstDirs, fsrcDirs, createEmptySrcDirs)
				if dirSyncErr != nil {
					f, errio := os.Create("./rlog.err")
					if errio != nil {
						return errio
					}
					defer f.Close()

					for _, errf := range errfs {
						s := fmt.Sprintf("Directory Sync Failture: %s\n", errf)
						f.WriteString(s)
					}
				}
			}
			if len(fsrcFileBases) > 0 && len(fsrcFileBases) == len(fdstFileNames) {
				errfs, fileSyncErr := operations.BatchCopyFiles(c, fdstFileBases, fsrcFileBases, fdstFileNames, fsrcFileNames)
				if fileSyncErr != nil {
					f, errio := os.OpenFile("./rlog.err",
						os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if errio != nil {
						return errio
					}
					defer f.Close()

					for _, errf := range errfs {
						s := fmt.Sprintf("File Sync Failure: %s\n", errf)
						f.WriteString(s)
					}
				}
			}

			if dirSyncErr != nil && fileSyncErr != nil {
				return fserrors.FatalError(errors.New("Both dirs and files have been failed to be copied"))
			}
			if dirSyncErr != nil {
				return dirSyncErr
			}
			if fileSyncErr != nil {
				return fileSyncErr
			}
			return nil
		})
	},
}
