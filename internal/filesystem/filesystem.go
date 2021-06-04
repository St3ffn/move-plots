// Package filesystem contains abstraction for filesystem and file related operations
package filesystem

import (
	"io"
	"os"
	"os/exec"
)

// Filesystem abstraction interface
type Filesystem interface {
	// Open a file
	Open(name string) (File, error)
	// Stat a file
	Stat(name string) (os.FileInfo, error)
}

// File abstraction interface
type File interface {
	io.Closer
	Readdir(n int) ([]os.FileInfo, error)
}

// LocalFs represents the local Filesystem
type LocalFs struct{}

// Open a file via os.Open
func (LocalFs) Open(name string) (File, error) {
	return os.Open(name)
}

// Stat a file via os.Stat
func (LocalFs) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

// Mover interface for moving a file
type Mover interface {
	Move(source, target string) error
}

// LocalMove represents the local file move operation
type LocalMove struct{}

// Move performs the mv os operation
func (LocalMove) Move(source, target string) error {
	return exec.Command("mv", source, target).Run()
}
