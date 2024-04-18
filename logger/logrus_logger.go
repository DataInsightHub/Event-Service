package logger

import "github.com/sirupsen/logrus"

type (
	LogrusLogger struct{}
)

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{}
}

func (*LogrusLogger) Info(msg string) {
	logrus.Infoln(msg)
}

func (*LogrusLogger) Warning(msg string) {
	logrus.Warnln(msg)
}

func (*LogrusLogger) Error(msg string) {
	logrus.Errorln(msg)
}

