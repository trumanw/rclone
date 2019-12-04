// Test Sftp filesystem interface

// +build !plan9

package sftp_test

import (
	"testing"

	"github.com/trumanw/rclone/backend/sftp"
	"github.com/trumanw/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSftp:",
		NilObject:  (*sftp.Object)(nil),
	})
}
