package main

import (
	"fmt"
	"io"
	"move-plots/internal/cli"
	"move-plots/internal/logging"
	"move-plots/internal/plot"
	"os"
)

var (
	GitVersion                   = "0.0.0"
	stdout     io.Writer         = os.Stdout
	stderr     io.Writer         = os.Stderr
	loggers    logging.LogFacade = logging.NewLoggers(stdout, stderr)
)

func main() {
	if err := run(); err != nil {
		loggers.E().Println(err.Error())
		os.Exit(1)
	}
}

// run the cli
func run() error {
	ctx, err := cli.RunCli(stdout, GitVersion)
	if err != nil {
		return err
	}
	if ctx.Done {
		return nil
	}

	if ctx.Verbose {
		loggers = logging.NewVerboseLoggers(stdout, stderr)
	}

	plots, err := plot.FindPlots(ctx.Source)
	if err != nil {
		return err
	}
	if len(plots) == 0 {
		loggers.I().Printf("no plots found in %s. nothing to do\n", ctx.Source)
		return nil
	}

	loggers.I().Printf("found %d plots in %s\n", len(plots), ctx.Source)

	for _, plotFilename := range plots {
		sourcePlot := plot.AbsoluteFilename(ctx.Source, plotFilename)

		loggers.V().Println("try to move", sourcePlot)

		targetDisk, err := plot.FindDisk(ctx.Reserved, ctx.Targets)
		if err != nil {
			return err
		}
		if targetDisk == nil {
			return fmt.Errorf("can not move plot %s. no disk space left", sourcePlot)
		}

		loggers.V().Printf("found %s with available capacity for %d plots\n", targetDisk.Path, targetDisk.PlotsLeft())

		targetPlot := plot.AbsoluteFilename(targetDisk.Path, plotFilename)

		loggers.V().Printf("try to move %s to %s\n", sourcePlot, targetPlot)

		if err := plot.MovePlot(ctx.Source, plotFilename, targetDisk); err != nil {
			return err
		}

		loggers.I().Printf("moved %s to %s\n", sourcePlot, targetPlot)
	}

	return nil
}
