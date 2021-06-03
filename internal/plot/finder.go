// Package plot contains plot related operations
package plot

import (
	"fmt"
	"github.com/St3ffn/plots-left/pkg/disk"
	"move-plots/internal/filesystem"
	"strings"
)

const plotSuffix = ".plot"

var Fs filesystem.Filesystem = filesystem.LocalFs{}

// FindPlots will find all *.plot files in given path. The function won't perform a recursive search.
func FindPlots(path string) (plots []string, err error) {
	dir, err := Fs.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %s", path, err.Error())
	}
	allFiles, err := dir.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("error reading %s: %s", path, err.Error())
	}
	err = dir.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing %s: %s", path, err.Error())
	}

	for _, file := range allFiles {
		if strings.HasSuffix(file.Name(), plotSuffix) {
			plots = append(plots, file.Name())
		}
	}
	return plots, nil
}

// FindDisk tries to find a disk with enough space for another plot
func FindDisk(reserved uint64, paths []string) (info *disk.PlotInfo, err error) {
	for _, path := range paths {
		d, err := disk.NewDisk(path)
		if err != nil {
			return nil, err
		}
		info = &disk.PlotInfo{
			Disk:     d,
			Reserved: reserved,
		}
		if info.PlotsLeft() > 0 {
			return info, nil
		}
	}
	return nil, nil
}
