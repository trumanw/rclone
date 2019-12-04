package serve

import (
	"errors"

	"github.com/trumanw/rclone/cmd"
	"github.com/trumanw/rclone/cmd/serve/dlna"
	"github.com/trumanw/rclone/cmd/serve/ftp"
	"github.com/trumanw/rclone/cmd/serve/http"
	"github.com/trumanw/rclone/cmd/serve/restic"
	"github.com/trumanw/rclone/cmd/serve/sftp"
	"github.com/trumanw/rclone/cmd/serve/webdav"
	"github.com/spf13/cobra"
)

func init() {
	Command.AddCommand(http.Command)
	if webdav.Command != nil {
		Command.AddCommand(webdav.Command)
	}
	if restic.Command != nil {
		Command.AddCommand(restic.Command)
	}
	if dlna.Command != nil {
		Command.AddCommand(dlna.Command)
	}
	if ftp.Command != nil {
		Command.AddCommand(ftp.Command)
	}
	if sftp.Command != nil {
		Command.AddCommand(sftp.Command)
	}
	cmd.Root.AddCommand(Command)
}

// Command definition for cobra
var Command = &cobra.Command{
	Use:   "serve <protocol> [opts] <remote>",
	Short: `Serve a remote over a protocol.`,
	Long: `rclone serve is used to serve a remote over a given protocol. This
command requires the use of a subcommand to specify the protocol, eg

    rclone serve http remote:

Each subcommand has its own options which you can see in their help.
`,
	RunE: func(command *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("serve requires a protocol, eg 'rclone serve http remote:'")
		}
		return errors.New("unknown protocol")
	},
}
