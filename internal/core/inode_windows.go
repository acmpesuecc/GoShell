//go:build windows
// +build windows

package core

import "os"

// getInode on Windows isn't reliably available; return false to indicate unsupported.
func getInode(info os.FileInfo) (uint64, bool) {
    return 0, false
}
