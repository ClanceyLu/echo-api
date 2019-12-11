package logger

import (
	"github.com/sirupsen/logrus"
)

// Init logrus
func Init() {
	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logrus.SetFormatter(formatter)
}
