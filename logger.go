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
	backendLeveled.SetLevel(level|logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backendLeveled, backendFormatter)

	return log
}

// Fatalf is equivalent to l.Critical followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	Log.Fatalf(format, args)
}

// Panicf is equivalent to l.Critical followed by a call to panic().
func Panicf(format string, args ...interface{}) {
	Log.Panicf(format, args)
}

// Criticalf logs a message using CRITICAL as log level.
func Criticalf(format string, args ...interface{}) {
	Log.Criticalf(format, args)
}

// Errorf logs a message using ERROR as log level.
func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args)
}

// Warningf logs a message using WARNING as log level.
func Warningf(format string, args ...interface{}) {
	Log.Warningf(format, args)
}

// Noticef logs a message using NOTICE as log level.
func Noticef(format string, args ...interface{}) {
	Log.Noticef(format, args)
}

// Infof logs a message using INFO as log level.
func Infof(format string, args ...interface{}) {
	Log.Infof(format, args)
}

// Debugf logs a message using DEBUG as log level.
func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args)
}

func Fatal(args ...interface{}) {
	Log.Fatal(args)
}

func Panic(args ...interface{}) {
	Log.Panic(args)
}

func Critical(args ...interface{}) {
	Log.Critical(args)
}

func Error(args ...interface{}) {
	Log.Error(args)
}

func Warning(args ...interface{}) {
	Log.Warning(args)
}

func Notice(args ...interface{}) {
	Log.Notice(args)
}

func Info(args ...interface{}) {
	Log.Info(args)
}

func Debug(args ...interface{}) {
	Log.Debug(args)
}
