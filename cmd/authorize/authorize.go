package authorize

import (
	"github.com/trumanw/rclone/cmd"
	"github.com/trumanw/rclone/fs/config"
	"github.com/trumanw/rclone/fs/config/flags"
	"github.com/spf13/cobra"
)

var (
	noAutoBrowser bool
)

func init() {
	cmd.Root.AddCommand(commandDefinition)
	cmdFlags := commandDefinition.Flags()
	flags.BoolVarP(cmdFlags, &noAutoBrowser, "auth-no-open-browser", "", false, "Do not automatically open auth link in default browser")
}

var commandDefinition = &cobra.Command{
	Use:   "authorize",
	Short: `Remote authorization.`,
	Long: `
Remote authorization. Used to authorize a remote or headless
rclone from a machine with a browser - use as instructed by
rclone config.

Use the --auth-no-open-browser to prevent rclone to open auth
link in default browser automatically.`,
	Run: func(command *cobra.Command, args []string) {
		cmd.CheckArgs(1, 3, command, args)
		config.Authorize(args, noAutoBrowser)
	},
}
