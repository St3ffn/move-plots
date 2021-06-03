package cli

import (
	"fmt"
	"move-plots/internal/filesystem"
)

// Validator interface to enforce certain rules
type Validator interface {
	// Enforce enforces the rule based on given path
	Enforce(path string) error
}

// IsDirectory is a validator to enforce a directory
type IsDirectory struct {
	filesystem.Filesystem
}

// Enforce enforces that given path is a directory
func (d IsDirectory) Enforce(path string) error {
	fileInfo, err := d.Stat(path)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}
	return nil
}
