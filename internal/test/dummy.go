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

type DummyFilesystem struct {
	Files      []string
	OpenErr    error
	CloseErr   error
	ReadDirErr error
	StatErr    error
	Directory  bool
}

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

func (t DummyFilesystem) Stat(name string) (fs.FileInfo, error) {
	if t.StatErr != nil {
		return nil, t.StatErr
	}
	return DummyFileInfo{
		name:      name,
		Directory: t.Directory,
	}, nil
}

type DummyFolder struct {
	files      []string
	closeErr   error
	readDirErr error
}

func (t DummyFolder) Close() error {
	return t.closeErr
}

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

type DummyFileInfo struct {
	name      string
	Directory bool
}

func (t DummyFileInfo) Name() string {
	return t.name
}

func (t DummyFileInfo) Size() int64 {
	panic("implement me")
}

func (t DummyFileInfo) Mode() fs.FileMode {
	panic("implement me")
}

func (t DummyFileInfo) ModTime() time.Time {
	panic("implement me")
}

func (t DummyFileInfo) IsDir() bool {
	return t.Directory
}

func (t DummyFileInfo) Sys() interface{} {
	panic("implement me")
}

type DummyStatfs struct {
	// pathPlotsLeft map with path as key and amount of available plots as value
	PathPlotsLeft map[string]uint64
	// err error to return when given path doesn't exist
	Err error
}

func (d DummyStatfs) Statfs(path string, stat *syscall.Statfs_t) (err error) {
	if available, exists := d.PathPlotsLeft[path]; exists {
		stat.Bsize = 1
		stat.Blocks = (available + 1) * disk.SizeOfPlot.Byte()
		stat.Bfree = available * disk.SizeOfPlot.Byte()
	}
	return d.Err
}