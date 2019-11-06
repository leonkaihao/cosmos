package log

import (
	"time"

	"github.com/sirupsen/logrus"
)

// Info ...
func Info(args ...interface{}) {
	logrus.Info(args...)
}

// Infof ...
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// Infoln ...
func Infoln(args ...interface{}) {
	logrus.Infoln(args...)
}

// Warn ...
func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

// Warnln ...
func Warnln(args ...interface{}) {
	logrus.Warnln(args...)
}

// Error ...
func Error(args ...interface{}) {
	logrus.Error(args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// Errorln ...
func Errorln(args ...interface{}) {
	logrus.Errorln(args...)
}

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
}
