# move-plots

[![Release](https://img.shields.io/github/v/release/St3ffn/move-plots)](https://github.com/St3ffn/move-plots/releases)
[![CI](https://github.com/St3ffn/move-plots/actions/workflows/ci.yml/badge.svg)](https://github.com/St3ffn/move-plots/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/st3ffn/move-plots)](/LICENSE)
[![GO](https://img.shields.io/github/go-mod/go-version/St3ffn/move-plots)](https://golang.org/)

![move](https://media.giphy.com/media/mMCNxRIZOtzny0S4bZ/giphy.gif)

Tiny CLI Tool to move chia plots from a source directory to a target directory with enough space left. 
The tool will work fine on all unix based systems (unix, linux, macos)

## Getting started

### Pre-Built Binaries

Pre-built binaries can be found on the [release page](https://github.com/St3ffn/move-plots/releases).
They are available for the following platforms:

- darwin-amd64 (64 Bit MacOS)
- linux-amd64 (64 Bit Linux)
- linux-arm64 (64Bit Linux for ARM)

### Building from Source

#### Pre-requisites

- Linux, MacOS or other Unix based System
- `git` installed
- `go 1.16` installed

### Installation 

Clone the repository

```shell
git clone https://github.com/St3ffn/move-plots.git
cd move-plots
```

Build the binary

```shell
go build
```

Now you are ready to go.

### Usage

```bash
# move all *.plot files from /path/source to a target directory with enough disk space left   
> move-plots /path/source /path/target1 /path/target2 /path/target3 

found 2 plots in /path/source
moved /path/source/marcus.plot to /path/target2/marcus.plot
moved /path/source/steffen.plot to /path/target2/steffen.plot
```

To specify the amount of plots to reserve use `--reserve` or `-r`
```bash
# move all *.plot files from /path/source to a target directory with enough disk space left   
# reserve space for 2 chia plots on each disk
> move-plots -r 2 /path/source /path/target1 /path/target2 /path/target3

...
```

To get more details use the verbose mode via `--verbose` or `-v`
```bash
# move all *.plot files from /path/source to a target directory with enough disk space left   
# reserve space for 2 chia plots on each disk
# verbose mode
> move-plots -V -r 2 /path/source /path/target1 /path/target2 /path/target3

INFO: 2021/06/04 15:47:19 found 2 plots in /path/source
INFO: 2021/06/04 15:47:19 try to move /path/source/marcus.plot
INFO: 2021/06/04 15:47:19 found /path/target2 with available capacity for 1 plots
INFO: 2021/06/04 15:47:19 try to move /path/source/marcus.plot to /path/target2/marcus.plot
INFO: 2021/06/04 15:47:19 moved /path/source/marcus.plot to /path/target2/marcus.plot
INFO: 2021/06/04 15:47:19 try to move /path/source/steffen.plot
INFO: 2021/06/04 15:47:19 found /path/target3 with available capacity for 5 plots
INFO: 2021/06/04 15:47:19 try to move /path/source/steffen.plot to /path/target3/steffen.plot
INFO: 2021/06/04 15:47:19 moved /path/source/steffen.plot to /path/target3/steffen.plot
```

To get the current version use `--version` or `-V`
```bash
> move-plots -V

move-plots version x.x.x
```

Call with `--help` or `-h` to see the help page
```bash
> ./move-plots -h
NAME:
   move-plots - move chia plots from source directory to a target directory with enough space left

USAGE:
   move-plots [-r RESERVE] [-v]  SOURCE_DIRECTORY TARGET_DIRECTORY ...
   move-plots -r 1 -v /source /plots/a /plots/b /plots/c

VERSION:
   0.1.0

DESCRIPTION:
   Tool will move each plot from source directory to a target directory with enough space left

GLOBAL OPTIONS:
   --reserve RESERVE, -r RESERVE  RESERVE. the amount of plots to reserve. (default: 0)
   --verbose, -v                  enable verbose mode. (default: false)
   --help, -h                     show help (default: false)
   --version, -V                  print version (default: false)

COPYRIGHT:
   GNU GPLv3
```

### Integration

The script can easily be integrated with cron. Simply open the users crontab via `crontab -e` and add the following line.

```shell
# run move-plots in verbose mode every 5 minutes and enforce that only one process is running (exclude grep and tail)
# all stdout and stderr will be redirected to /tmp/move-plots.log
*/5 * * * * [ $(ps aux | grep move-plots| grep -v grep | grep -v tail | wc -l) -eq 0 ] && /PATH/TO/move-plots -v /plots/source /target/one /target/another >>/tmp/move-plots.log 2>&1
```

## Limitations and Restrictions

- by default the capacity of one plot is reserved for each target disk (can be disabled via `-r 0`)
- plots are moved one after another (sequentially)
- before a single plot gets moved each target disk is evaluated for plot capacity
- the actual move operation is performed by the `mv` command (don't wanted to reinvent the wheel)
- in case there is no disk space left for further plots, the script will error `...no disk space left`
- windows is not supported and there is also no plan to support it

## Kind gestures

If you like the tool, and you are open for a kind gesture. Thanks in advance. 

- XCH Address: xch18s8r9v4kpwdx2y8jks5ma4g2rmff0h9dtr5nkc6zmnk5kj6v0faqer6k9v

