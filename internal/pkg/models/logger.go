package models

import "github.com/sirupsen/logrus"

type Logger struct {
	Logger *logrus.Entry
}

type LoggerInterface interface {
	LogInfo(data interface{})
	LogError(data interface{})
}

func (l *Logger) LogInfo(data interface{}) {
	l.Logger.Info(data)
}

func (l *Logger) LogError(data interface{}) {
	l.Logger.Error(data)
}
