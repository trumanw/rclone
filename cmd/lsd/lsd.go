package lsd

import (
	"context"
	"os"

	"github.com/trumanw/rclone/cmd"
	"github.com/trumanw/rclone/cmd/ls/lshelp"
	"github.com/trumanw/rclone/fs"
	"github.com/trumanw/rclone/fs/config/flags"
	"github.com/trumanw/rclone/fs/operations"
	"github.com/spf13/cobra"
)

var (
	recurse bool
)

func init() {
	cmd.Root.AddCommand(commandDefinition)
	cmdFlags := commandDefinition.Flags()
	flags.BoolVarP(cmdFlags, &recurse, "recursive", "R", false, "Recurse into the listing.")
}

var commandDefinition = &cobra.Command{
	Use:   "lsd remote:path",
	Short: `List all directories/containers/buckets in the path.`,
	Long: `
Lists the directories in the source path to standard output. Does not
recurse by default.  Use the -R flag to recurse.

This command lists the total size of the directory (if known, -1 if
not), the modification time (if known, the current time if not), the
number of objects in the directory (if known, -1 if not) and the name
of the directory, Eg

    $ rclone lsd swift:
          494000 2018-04-26 08:43:20     10000 10000files
              65 2018-04-26 08:43:20         1 1File

Or

    $ rclone lsd drive:test
              -1 2016-10-17 17:41:53        -1 1000files
              -1 2017-01-03 14:40:54        -1 2500files
              -1 2017-07-08 14:39:28        -1 4000files

If you just want the directory names use "rclone lsf --dirs-only".

` + lshelp.Help,
	Run: func(command *cobra.Command, args []string) {
		cmd.CheckArgs(1, 1, command, args)
		if recurse {
			fs.Config.MaxDepth = 0
		}
		fsrc := cmd.NewFsSrc(args)
		cmd.Run(false, false, command, func() error {
			return operations.ListDir(context.Background(), fsrc, os.Stdout)
		})
	},
}
