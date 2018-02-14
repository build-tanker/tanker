package logger

import (
	"io"
	"log"

	"github.com/gojektech/tanker/pkg/config"
	"github.com/sirupsen/logrus"
)

// Logger - inteface for logrus
type Logger interface {
	Fatalln(args ...interface{})
	Errorln(args ...interface{})
	Debugln(args ...interface{})
	Infoln(args ...interface{})
}

type logger struct {
	l *logrus.Logger
}

// NewLogger - create a new logrus logger
func NewLogger(config *config.Config, w io.Writer) Logger {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		log.Fatalf(err.Error())
	}

	l := &logrus.Logger{
		Out:       w,
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
		Formatter: &logrus.TextFormatter{},
	}

	return &logger{
		l: l,
	}
}

func (l *logger) Fatalln(args ...interface{}) {
	l.l.Fatalln(args...)
}

func (l *logger) Errorln(args ...interface{}) {
	l.l.Errorln(args...)
}

func (l *logger) Debugln(args ...interface{}) {
	l.l.Debugln(args...)
}

func (l *logger) Infoln(args ...interface{}) {
	l.l.Infoln(args...)
}
