// Package logging contains log specific functionality
package logging

import (
	"io"
	"log"
)

// InfoPrefix for info log level prefix
const InfoPrefix = "INFO: "

// ErrorPrefix for error log level prefix
const ErrorPrefix = "ERROR: "

// VerbosePrefix for verbose log level prefix
const VerbosePrefix = "INFO: " // same as info on purpose

// LogFacade acts as facade for the loggers of the different log level
type LogFacade interface {
	// I gets the info level logger
	I() *log.Logger
	// E gets the error level logger
	E() *log.Logger
	// V gets the verbose level logger
	V() *log.Logger
}

// Loggers stores the loggers of the different log levels
type Loggers struct {
	info    Logger
	error   Logger
	verbose Logger
}

// NewLoggers creates a new facade for the different log levels
func NewLoggers(stdout, stderr io.Writer) *Loggers {
	return &Loggers{
		info:    newLog(stdout),
		error:   newLog(stderr),
		verbose: newLog(io.Discard),
	}
}

// NewVerboseLoggers creates a more verbose facade for the different log levels
func NewVerboseLoggers(stdout, stderr io.Writer) *Loggers {
	return &Loggers{
		info:    newVerboseLogger(stdout, InfoPrefix),
		error:   newVerboseLogger(stderr, ErrorPrefix),
		verbose: newVerboseLogger(stdout, VerbosePrefix),
	}
}

// I get the info level logger
func (l Loggers) I() *log.Logger {
	return l.info.Logger
}

// E get the error level logger
func (l Loggers) E() *log.Logger {
	return l.error.Logger
}

// V get the verbose level logger
func (l Loggers) V() *log.Logger {
	return l.verbose.Logger
}

// Logger is a wrapper around log.Logger
type Logger struct {
	*log.Logger
	// verbose indicates if verbose mode is enabled
	verbose bool
}

// newLog creates a new io.Logger for the given writer
func newLog(writer io.Writer) Logger {
	return Logger{
		Logger:  log.New(writer, "", 0),
		verbose: false,
	}
}

// newVerboseLogger creates a new verbose io.Logger for the given writer and prefix.
func newVerboseLogger(writer io.Writer, prefix string) Logger {
	return Logger{
		Logger:  log.New(writer, prefix, log.Ldate|log.Ltime),
		verbose: true,
	}
}
