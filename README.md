# slog

A simple logger for Go with no dependencies

I know there are many go loggers out there that offer various logging features such as file rotation, granular verbosity settings, colored and JSON output, etc.

_Slog is not one of them._

Slog lets you hide or show debug logs as well as provides a simpler way to log messages with Warning and Error prefixes for consistency.

Also provided are a few simple methods for handling returned `error` variables, logging them out and optionally panicing or fatally exiting.

## Documentation
    package slog // import "git.iamthefij.com/iamthefij/slog"

    Package slog is a super simple logger that allows a few convenience methods
    for handling debug vs warning/error logs. It also adds a few conveniences
    for handling errors.

    VARIABLES

    var (
    	// DebugLevel indicates if we should log at the debug level
    	DebugLevel = true

    	// LoggerInfo is the slog Info logger
    	LoggerInfo = log.New(os.Stderr, "INFO    ", defaultFlags)
    	// LoggerWarning is the slog Warning logger
    	LoggerWarning = log.New(os.Stderr, "WARNING ", defaultFlags)
    	// LoggerError is the slog Error logger
    	LoggerError = log.New(os.Stderr, "ERROR   ", defaultFlags)
    	// LoggerDebug is the slog Debug logger
    	LoggerDebug = log.New(os.Stderr, "DEBUG   ", defaultFlags)
    )

    FUNCTIONS

    func Debugf(format string, v ...interface{})
        Debugf will log with a DEBUG prefix if DebugLevel is se

    func Errorf(format string, v ...interface{})
        Errorf will log with a ERROR prefix

    func Fatalf(format string, v ...interface{})
        Fatalf will log with a ERROR prefix followed by exit(1)

    func Infof(format string, v ...interface{})
        Infof formats logs with an INFO prefix

    func Logf(format string, v ...interface{})
        Logf formats logs directly to the main logger

    func OnErrFatalf(err error, format string, v ...interface{})
        OnErrFatalf if error provided, will log out details of an error and exi

    func OnErrPanicf(err error, format string, v ...interface{})
        OnErrPanicf if error provided, will log out details of an error and exi

    func OnErrWarnf(err error, format string, v ...interface{})
        OnErrWarnf if error provided, will provide a warning if an error is provided

    func Panicf(format string, v ...interface{})
        Panicf will log with a ERROR prefix followed by panic()

    func SetFlags(flag int)
        SetFlags allows changing the logger flags using flags found in `log`

    func SetOutput(w io.Writer)
        SetOutput allows changing the output of all loggers

    func Warningf(format string, v ...interface{})
        Warningf will log with a WARNING prefix
