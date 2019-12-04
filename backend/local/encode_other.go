//+build !windows,!darwin

package local

import (
	"github.com/trumanw/rclone/fs/encodings"
)

const enc = encodings.LocalUnix
