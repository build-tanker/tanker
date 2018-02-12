package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/sudhanshuraheja/tanker/pkg/config"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(config *config.Config) Logger {
	level, err := logrus.ParseLevel(config.LogLevel())
	if err != nil {
		log.Fatalf(err.Error())
	}

	return Logger{
		&logrus.Logger{
			Out:       os.Stdout,
			Hooks:     make(logrus.LevelHooks),
			Level:     level,
			Formatter: &logrus.JSONFormatter{},
		},
	}
}
