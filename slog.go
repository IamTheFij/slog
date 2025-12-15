// Package slog is a super simple logger that allows a few convenience methods
// for handling debug vs warning/error logs. It also adds a few conveniences for
// handling errors.
package slog

import (
	"io"
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

	allLoggers = []*log.Logger{
		LoggerInfo,
		LoggerWarning,
		LoggerError,
		LoggerDebug,
	}
)

// SetFlags allows changing the logger flags using flags found in `log`
func SetFlags(flag int) {
	for _, logger := range allLoggers {
		logger.SetFlags(flag)
	}
}

// SetOutput allows changing the output of all loggers
func SetOutput(w io.Writer) {
	for _, logger := range allLoggers {
		logger.SetOutput(w)
	}
}

// Logf formats logs directly to the main logger
func Logf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Debug will log with a DEBUG prefix if DebugLevel is set
func Debug(v ...interface{}) {
	if !DebugLevel {
		return
	}

	LoggerDebug.Print(v...)
}

// Debugln will log with a DEBUG prefix if DebugLevel is set
func Debugln(v ...interface{}) {
	if !DebugLevel {
		return
	}

	LoggerDebug.Println(v...)
}

// Debugf will log with a DEBUG prefix if DebugLevel is set
func Debugf(format string, v ...interface{}) {
	if !DebugLevel {
		return
	}

	LoggerDebug.Printf(format, v...)
}

// Print formats logs with an INFO prefix
func Print(v ...interface{}) {
	LoggerInfo.Print(v...)
}

// Println formats logs with an INFO prefix
func Println(v ...interface{}) {
	LoggerInfo.Println(v...)
}

// Printf formats logs with an INFO prefix
func Printf(format string, v ...interface{}) {
	LoggerInfo.Printf(format, v...)
}

// Info formats logs with an INFO prefix
func Info(v ...interface{}) {
	LoggerInfo.Print(v...)
}

// Infoln formats logs with an INFO prefix
func Infoln(v ...interface{}) {
	LoggerInfo.Println(v...)
}

// Infof formats logs with an INFO prefix
func Infof(format string, v ...interface{}) {
	LoggerInfo.Printf(format, v...)
}

// Warning will log with a WARNING prefix
func Warning(v ...interface{}) {
	LoggerWarning.Print(v...)
}

// Warningln will log with a WARNING prefix
func Warningln(v ...interface{}) {
	LoggerWarning.Println(v...)
}

// Warningf will log with a WARNING prefix
func Warningf(format string, v ...interface{}) {
	LoggerWarning.Printf(format, v...)
}

// Error will log with a ERROR prefix
func Error(v ...interface{}) {
	LoggerError.Print(v...)
}

// Errorln will log with a ERROR prefix
func Errorln(v ...interface{}) {
	LoggerError.Println(v...)
}

// Errorf will log with a ERROR prefix
func Errorf(format string, v ...interface{}) {
	LoggerError.Printf(format, v...)
}

// Fatal will log with a ERROR prefix followed by exit(1)
func Fatal(v ...interface{}) {
	LoggerError.Fatal(v...)
}

// Fatalln will log with a ERROR prefix followed by exit(1)
func Fatalln(v ...interface{}) {
	LoggerError.Fatalln(v...)
}

// Fatalf will log with a ERROR prefix followed by exit(1)
func Fatalf(format string, v ...interface{}) {
	LoggerError.Fatalf(format, v...)
}

// Panic will log with a ERROR prefix followed by panic()
func Panic(v ...interface{}) {
	LoggerError.Panic(v...)
}

// Panicln will log with a ERROR prefix followed by panic()
func Panicln(v ...interface{}) {
	LoggerError.Panicln(v...)
}

// Panicf will log with a ERROR prefix followed by panic()
func Panicf(format string, v ...interface{}) {
	LoggerError.Panicf(format, v...)
}

// OnErrWarnf if error provided, will provide a warning if an error is provided
func OnErrWarnf(err error, format string, v ...interface{}) {
	if err != nil {
		LoggerWarning.Printf(format, v...)
		LoggerError.Print(err)
	}
}

// OnErrFatalf if error provided, will log out details of an error and exit
func OnErrFatalf(err error, format string, v ...interface{}) {
	if err != nil {
		LoggerError.Printf(format, v...)
		LoggerError.Fatal(err)
	}
}

// OnErrPanicf if error provided, will log out details of an error and exit
func OnErrPanicf(err error, format string, v ...interface{}) {
	if err != nil {
		LoggerError.Printf(format, v...)
		LoggerError.Panic(err)
	}
}
