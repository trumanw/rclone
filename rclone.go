// Sync files and directories to and from local and remote object stores
//
// Nick Craig-Wood <nick@craig-wood.com>
package main

import (
	// _ "github.com/trumanw/rclone/backend/all" // import all backends
	_ "github.com/trumanw/rclone/backend/all" // import all backends from trumanw's fork
	"github.com/trumanw/rclone/cmd"
	_ "github.com/trumanw/rclone/cmd/all"    // import all commands
	_ "github.com/trumanw/rclone/lib/plugin" // import plugins
)

func main() {
	cmd.Main()
}
