// Package cli contains the command line interface, defines related parameters and validates them
package cli

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"move-plots/internal/filesystem"
	"os"
	"strconv"
)

var (
	Args = os.Args
	// Reserved defines the default amount of plots to reserve
	Reserved uint64 = 0
	// Validation defines the validator to be used
	Validation Validator = IsDirectory{filesystem.LocalFs{}}
)

// Context describes the environment of the tool execution
type Context struct {
	// Reserved represents the amount of plots to be reserved
	Reserved uint64
	// Source directory containing plots
	Source string
	// Target directories to choose from
	Targets []string
	// Verbose indicates verbose mode
	Verbose bool
	// Done indicates that we are done (--help, --version...)
	Done bool
}

// RunCli starts the cli which includes validation of parameters.
func RunCli(writer io.Writer, version string) (*Context, error) {
	var source string
	var targets []string
	var verbose, done bool

	cli.HelpFlag = &cli.BoolFlag{
		Name:        "help",
		Aliases:     []string{"h"},
		Usage:       "show help",
		Destination: &done,
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "print version",
	}
	cli.VersionPrinter = func(c *cli.Context) {
		_, _ = fmt.Fprintf(c.App.Writer, "%s version %s\n", c.App.Name, c.App.Version)
		done = true
	}

	app := &cli.App{
		Name:                 "move-plots",
		Usage:                "move chia plots from source directory to a target directory with enough space left",
		UsageText:            "move-plots [-r RESERVE] [-v]  SOURCE_DIRECTORY TARGET_DIRECTORY ...\n\t move-plots -r 0 -v /source /plots/a /plots/b /plots/c",
		Description:          "Tool will move each plot from source directory to a target directory with enough space left",
		EnableBashCompletion: true,
		HideHelpCommand:      true,
		Version:              version,
		Flags: []cli.Flag{
			&cli.Uint64Flag{
				Name:        "reserve",
				Aliases:     []string{"r"},
				Required:    false,
				Value:       Reserved,
				DefaultText: strconv.FormatUint(Reserved, 10),
				Usage:       "`RESERVE`. the amount of plots to reserve.",
				Destination: &Reserved,
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Required:    false,
				Value:       false,
				Usage:       "enable verbose mode.",
				Destination: &verbose,
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return errors.New("SOURCE_DIRECTORY and TARGET_DIRECTORY missing")
			}
			if c.NArg() < 2 {
				return errors.New("TARGET_DIRECTORY missing")
			}

			for _, path := range c.Args().Slice() {
				if err := Validation.Enforce(path); err != nil {
					return err
				}
			}

			source = c.Args().First()
			targets = c.Args().Slice()[1:]

			return nil
		},
		Copyright: "GNU GPLv3",
	}
	app.Writer = writer

	err := app.Run(Args)
	if err != nil {
		return nil, err
	}

	return &Context{
		Reserved: Reserved,
		Source:   source,
		Targets:  targets,
		Verbose:  verbose,
		Done:     done,
	}, nil
}
