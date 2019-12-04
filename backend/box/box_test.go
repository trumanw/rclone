// Test Box filesystem interface
package box_test

import (
	"testing"

	"github.com/trumanw/rclone/backend/box"
	"github.com/trumanw/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestBox:",
		NilObject:  (*box.Object)(nil),
	})
}
