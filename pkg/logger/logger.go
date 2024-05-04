package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	})

	if logLevel := os.Getenv("LOG_LEVEL"); logLevel == "" {
		logger.SetLevel(logrus.ErrorLevel)
	} else {
		logLevel, err := logrus.ParseLevel(logLevel)
		if err != nil {
			logger.Errorf("invalid log level '%s', defaulting to 'error'", logLevel)
			logger.SetLevel(logrus.ErrorLevel)
		} else {
			logger.SetLevel(logLevel)
		}

	}

	return logger
}
