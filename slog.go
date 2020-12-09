// Package slog is a super simple logger that allows a few convenience methods
// for handling debug vs warning/error logs. It also adds a few conveniences for
// handling errors.
package slog

import (
	"log"
	"os"
)

var (
	// DebugLevel indicates if we should log at the debug level
	DebugLevel = true

	// Default set of flags to use
	defaultFlags = log.LstdFlags | log.Lmsgprefix

	// Loggers for various levels. Prefixes are padded to align logged content

	// LoggerInfo is the slog Info logger
	LoggerInfo = log.New(os.Stderr, "INFO    ", defaultFlags)
	// LoggerWarning is the slog Warning logger
	LoggerWarning = log.New(os.Stderr, "WARNING ", defaultFlags)
	// LoggerError is the slog Error logger
	LoggerError = log.New(os.Stderr, "ERROR   ", defaultFlags)
	// LoggerDebug is the slog Debug logger
	LoggerDebug = log.New(os.Stderr, "DEBUG   ", defaultFlags)
)

// SetFlags allows changing the logger flags using flags found in `log`
func SetFlags(flag int) {
	for _, logger := range []*log.Logger{
		LoggerInfo,
		LoggerWarning,
		LoggerError,
		LoggerDebug,
	} {
		logger.SetFlags(flag)
	}
}

// Log formats logs directly to the main logger
func Log(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Info formats logs with an INFO prefix
func Info(format string, v ...interface{}) {
	LoggerInfo.Printf(format, v...)
}

// Warning will log with a WARNING prefix
func Warning(format string, v ...interface{}) {
	LoggerWarning.Printf(format, v...)
}

// Error will log with a ERROR prefix
func Error(format string, v ...interface{}) {
	LoggerError.Printf(format, v...)
}

// Debug will log with a DEBUG prefix if DebugLevel is set
func Debug(format string, v ...interface{}) {
	if !DebugLevel {
		return
	}
	LoggerDebug.Printf(format, v...)
}

// WarnOnErr if error provided, will provide a warning if an error is provided
func WarnOnErr(err error, format string, v ...interface{}) {
	if err != nil {
		LoggerWarning.Printf(format, v...)
		LoggerError.Print(err)
	}
}

// FatalOnErr if error provided, will log out details of an error and exit
func FatalOnErr(err error, format string, v ...interface{}) {
	if err != nil {
		LoggerError.Printf(format, v...)
		LoggerError.Fatal(err)
	}
}

// PanicOnErr if error provided, will log out details of an error and exit
func PanicOnErr(err error, format string, v ...interface{}) {
	if err != nil {
		LoggerError.Printf(format, v...)
		LoggerError.Panic(err)
	}
}
