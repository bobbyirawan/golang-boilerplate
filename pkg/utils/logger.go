package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

type (
	Logger interface {
		Info()
	}
	implLogger struct {
		logger *logrus.Logger
		time   TimeService
	}
)

func NewLogger(time TimeService) Logger {
	logger := logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return &implLogger{
		logger: logger,
		time:   time,
	}
}

func (impl *implLogger) Info() {
	impl.logger.WithTime(impl.time.Indonesia()).WithFields(logrus.Fields{
		"requestId": GenerateUUID(),
	}).Info("hello testing")
}

func (impl *implLogger) Warn() {

}
