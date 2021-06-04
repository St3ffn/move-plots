// Package test contains test related code
package test

import (
	"github.com/St3ffn/plots-left/pkg/disk"
	"io/fs"
	"move-plots/internal/filesystem"
	"os"
	"syscall"
	"time"
)

// DummyFilesystem is a testing filesystem
type DummyFilesystem struct {
	Files      []string
	OpenErr    error
	CloseErr   error
	ReadDirErr error
	StatErr    error
	Directory  bool
}

// Open returns a DummyFolder or error if defined
func (t DummyFilesystem) Open(_ string) (filesystem.File, error) {
	if t.OpenErr != nil {
		return nil, t.OpenErr
	}

	return DummyFolder{
		files:      t.Files,
		closeErr:   t.CloseErr,
		readDirErr: t.ReadDirErr,
	}, nil
}

// Stat returns DummyFileInfo or error if defined
func (t DummyFilesystem) Stat(name string) (fs.FileInfo, error) {
	if t.StatErr != nil {
		return nil, t.StatErr
	}
	return DummyFileInfo{
		name:      name,
		Directory: t.Directory,
	}, nil
}

// DummyFolder which consists of files. Struct can also contain error for closing the directory or reading the directory.
type DummyFolder struct {
	files      []string
	closeErr   error
	readDirErr error
}

// Close return the closing error
func (t DummyFolder) Close() error {
	return t.closeErr
}

// Readdir returns the list of files in the directory or error if defined
func (t DummyFolder) Readdir(n int) (infos []os.FileInfo, err error) {
	if t.readDirErr != nil {
		return nil, t.readDirErr
	}

	for _, filename := range t.files {
		infos = append(infos, DummyFileInfo{
			name: filename,
		})
	}
	return infos, nil
}

// DummyFileInfo is a testing file info
type DummyFileInfo struct {
	name      string
	Directory bool
}

// Name the name of the file
func (t DummyFileInfo) Name() string {
	return t.name
}

// Size is not implemented
func (t DummyFileInfo) Size() int64 {
	panic("implement me")
}

// Mode is not implemented
func (t DummyFileInfo) Mode() fs.FileMode {
	panic("implement me")
}

// ModTime is not implemented
func (t DummyFileInfo) ModTime() time.Time {
	panic("implement me")
}

// IsDir is not implemented
func (t DummyFileInfo) IsDir() bool {
	return t.Directory
}

// Sys is not implemented
func (t DummyFileInfo) Sys() interface{} {
	panic("implement me")
}

// DummyStatfs testing statfs operation representing certain amount of plots in paths
type DummyStatfs struct {
	// pathPlotsLeft map with path as key and amount of available plots as value
	PathPlotsLeft map[string]uint64
	// err error to return when given path doesn't exist
	Err error
}

// Statfs returns the block size information for the defined amount of plots or error if defined
func (d DummyStatfs) Statfs(path string, stat *syscall.Statfs_t) (err error) {
	if available, exists := d.PathPlotsLeft[path]; exists {
		stat.Bsize = 1
		stat.Blocks = (available + 1) * disk.SizeOfPlot.Byte()
		stat.Bfree = available * disk.SizeOfPlot.Byte()
	}
	return d.Err
}

// DummyMover is a testing mover
type DummyMover struct {
	// TargetFileError represents a map with the targets (path + filename) and error to return for move operation
	TargetFileError map[string]error
}

// Move return error for given target if defined
func (d DummyMover) Move(_, target string) error {
	return d.TargetFileError[target]
}
