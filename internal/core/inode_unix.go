//go:build linux || darwin || freebsd || netbsd || openbsd || dragonfly
// +build linux darwin freebsd netbsd openbsd dragonfly

package core

import (
    "os"
    "syscall"
)

// getInode returns the inode number for Unix systems.
func getInode(info os.FileInfo) (uint64, bool) {
    if stat, ok := info.Sys().(*syscall.Stat_t); ok {
        return uint64(stat.Ino), true
    }
    return 0, false
}
