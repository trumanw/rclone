// Test Cache filesystem interface

// +build !plan9

package cache_test

import (
	"testing"

	"github.com/trumanw/rclone/backend/cache"
	_ "github.com/trumanw/rclone/backend/local"
	"github.com/trumanw/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName:                   "TestCache:",
		NilObject:                    (*cache.Object)(nil),
		UnimplementableFsMethods:     []string{"PublicLink", "MergeDirs", "OpenWriterAt"},
		UnimplementableObjectMethods: []string{"MimeType", "ID", "GetTier", "SetTier"},
		SkipInvalidUTF8:              true, // invalid UTF-8 confuses the cache
	})
}
