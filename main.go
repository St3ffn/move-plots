package main

import (
	"fmt"
	"github.com/St3ffn/plots-left/pkg/disk"
	"io"
	"move-plots/internal/cli"
	"os"
)

var (
	GitVersion           = "0.0.0"
	stderr     io.Writer = os.Stderr
	stdout     io.Writer = os.Stdout
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(stderr, "%s\n", err)
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
	_, err = disk.NewDisk(ctx.Source)
	if err != nil {
		return err
	}
	return nil
}
