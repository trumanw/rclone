// Test Union filesystem interface
package union_test

import (
	"testing"

	_ "github.com/trumanw/rclone/backend/local"
	"github.com/trumanw/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName:  "TestUnion:",
		NilObject:   nil,
		SkipFsMatch: true,
	})
}
