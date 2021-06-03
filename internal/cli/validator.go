package cli

import (
	"fmt"
	"io/fs"
)

type Validator interface {
	Enforce(path string) error
}

type IsDirectory struct {
	stat func (name string) (fs.FileInfo, error)
}

func (d IsDirectory) Enforce(path string) error {
	fileInfo, err := d.stat(path)
	if err != nil {
		return err
	}
	if ! fileInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}
	return nil
}
