package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var l Logger

func SetLogger(logger Logger) {
	l = logger
}

type Logger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

func Error(args ...interface{})                 { l.Error(args...) }
func Errorf(format string, args ...interface{}) { l.Errorf(format, args...) }
func Debug(args ...interface{})                 { l.Debug(args...) }
func Debugf(format string, args ...interface{}) { l.Debugf(format, args...) }
func Warn(args ...interface{})                  { l.Warn(args...) }
func Warnf(format string, args ...interface{})  { l.Warnf(format, args...) }
func Fatal(args ...interface{})                 { l.Fatal(args...) }
func Fatalf(format string, args ...interface{}) { l.Fatalf(format, args...) }

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	l = log.StandardLogger()
}
