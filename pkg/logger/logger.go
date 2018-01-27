package logger

import (
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/sudhanshuraheja/tanker/pkg/config"
)

var logger *logrus.Logger

// Init : setup the logger
func Init() {
	level, err := logrus.ParseLevel(config.LogLevel())
	if err != nil {
		log.Fatalf(err.Error())
	}

	logger = &logrus.Logger{
		Out:       os.Stdout,
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
		Formatter: &logrus.TextFormatter{},
	}
}

// AddHook : add a logrus hook
func AddHook(hook logrus.Hook) {
	logger.Hooks.Add(hook)
}

// Debug : log a debug error
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf : log a formatted debug error
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Debugln : log a debug error with new line
func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

// Error : log an error
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf : log a formatted error
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Errorln : log an error with new line
func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}

// Errorrf : log a formatted request error
func Errorrf(r *http.Request, format string, args ...interface{}) {
	httpRequestLogEntry(r).Errorf(format, args...)
}

// Info : log an informational error
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof : log an formatted informational error
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Infoln : log an informational error with new line
func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

// Inforf : log an informational request error
func Inforf(r *http.Request, format string, args ...interface{}) {
	httpRequestLogEntry(r).Infof(format, args...)
}

// Warn : log a warning
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf : log a formatted warning
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Warnln : log a warning with a new line
func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

// Fatal : log a warning
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf : log a formatted warning
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Fatalln : log a warning with a new line
func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

// WithField : log with one field
func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

// WithFields : log with multiple fields
func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

func httpRequestLogEntry(r *http.Request) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"RequestMethod": r.Method,
		"Host":          r.Host,
		"Path":          r.URL.Path,
	})
}
