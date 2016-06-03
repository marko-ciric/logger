package log

import (
	"github.com/op/go-logging"
	"os"
)

// Log is a default reference
var Log *logging.Logger

func init() {
	Log = GetLog(logging.ERROR)
}

// GetLog is a reference to a logging object
func GetLog(level logging.Level) *logging.Logger {

	log := logging.MustGetLogger("default")

	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	backend := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backendFormatter := logging.NewBackendFormatter(backend, format)

	// Only errors and more severe messages should be sent to backend
	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(level, "")

	// Set the backends to be used.
	logging.SetBackend(backendLeveled, backendFormatter)
	return log
}

// Fatal is equivalent to l.Critical followed by a call to os.Exit(1).
func Fatal(format string, args ...interface{}) {
	Log.Fatalf(format, args)
}

// Panic is equivalent to l.Critical followed by a call to panic().
func Panic(format string, args ...interface{}) {
	Log.Panicf(format, args)
}

// Critical logs a message using CRITICAL as log level.
func Critical(format string, args ...interface{}) {
	Log.Criticalf(format, args)
}

// Error logs a message using ERROR as log level.
func Error(format string, args ...interface{}) {
	Log.Errorf(format, args)
}

// Warning logs a message using WARNING as log level.
func Warning(format string, args ...interface{}) {
	Log.Warningf(format, args)
}

// Notice logs a message using NOTICE as log level.
func Notice(format string, args ...interface{}) {
	Log.Noticef(format, args)
}

// Info logs a message using INFO as log level.
func Info(format string, args ...interface{}) {
	Log.Infof(format, args)
}

// Debug logs a message using DEBUG as log level.
func Debug(format string, args ...interface{}) {
	Log.Debugf(format, args)
}
