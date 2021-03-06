// Test the Chunker filesystem interface
package chunker_test

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/trumanw/rclone/backend/all" // for integration tests
	"github.com/trumanw/rclone/backend/chunker"
	"github.com/trumanw/rclone/fstest"
	"github.com/trumanw/rclone/fstest/fstests"
)

// Command line flags
var (
	// Invalid characters are not supported by some remotes, eg. Mailru.
	// We enable testing with invalid characters when -remote is not set, so
	// chunker overlays a local directory, but invalid characters are disabled
	// by default when -remote is set, eg. when test_all runs backend tests.
	// You can still test with invalid characters using the below flag.
	UseBadChars = flag.Bool("bad-chars", false, "Set to test bad characters in file names when -remote is set")
)

// TestIntegration runs integration tests against a concrete remote
// set by the -remote flag. If the flag is not set, it creates a
// dynamic chunker overlay wrapping a local temporary directory.
func TestIntegration(t *testing.T) {
	opt := fstests.Opt{
		RemoteName:               *fstest.RemoteName,
		NilObject:                (*chunker.Object)(nil),
		SkipBadWindowsCharacters: !*UseBadChars,
		UnimplementableObjectMethods: []string{
			"MimeType",
			"GetTier",
			"SetTier",
		},
		UnimplementableFsMethods: []string{
			"PublicLink",
			"OpenWriterAt",
			"MergeDirs",
			"DirCacheFlush",
			"UserInfo",
			"Disconnect",
		},
	}
	if *fstest.RemoteName == "" {
		name := "TestChunker"
		opt.RemoteName = name + ":"
		tempDir := filepath.Join(os.TempDir(), "rclone-chunker-test-standard")
		opt.ExtraConfig = []fstests.ExtraConfigItem{
			{Name: name, Key: "type", Value: "chunker"},
			{Name: name, Key: "remote", Value: tempDir},
		}
	}
	fstests.Run(t, &opt)
}
