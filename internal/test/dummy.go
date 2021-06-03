// Package test contains test related code
package test

import (
	"io/fs"
	"move-plots/internal/filesystem"
	"os"
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
