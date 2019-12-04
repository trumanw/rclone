// Test Local filesystem interface
package local_test

import (
	"testing"

	"github.com/trumanw/rclone/backend/local"
	"github.com/trumanw/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "",
		NilObject:  (*local.Object)(nil),
	})
}
