package plot

import (
	"fmt"
	"github.com/St3ffn/plots-left/pkg/disk"
	"move-plots/internal/filesystem"
)

var Mover filesystem.Mover = filesystem.LocalMove{}

func MovePlot(sourcePath, sourceFilename string, targetDisk *disk.PlotInfo) error {
	sourcePlot := AbsoluteFilename(sourcePath, sourceFilename)
	targetTmpPlot := AbsoluteFilename(targetDisk.Path, sourceFilename+"-mv")
	if err := Mover.Move(sourcePlot, targetTmpPlot); err != nil {
		return fmt.Errorf("can not move %s to %s: %s", sourcePlot, targetTmpPlot, err.Error())
	}

	targetPlot := AbsoluteFilename(targetDisk.Path, sourceFilename)
	if err := Mover.Move(targetTmpPlot, targetPlot); err != nil {
		return fmt.Errorf("can not move %s to %s: %s", targetTmpPlot, targetPlot, err.Error())
	}
	return nil
}

// AbsoluteFilename concat path with filename via "/" separator
func AbsoluteFilename(path, filename string) string {
	return path + "/" + filename
}
