package logger

import (
	"github.com/RichardKnop/logging"
	"github.com/sirupsen/logrus"
)

type debugLogger struct{}

func NewDebugLogger() logging.LoggerInterface {
	return &debugLogger{}
}

func (debugLogger) Print(args ...interface{}) {
	logrus.Debug(args...)
}

func (debugLogger) Printf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func (debugLogger) Println(args ...interface{}) {
	logrus.Debugln(args...)
}

func (debugLogger) Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func (debugLogger) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func (debugLogger) Fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}

func (debugLogger) Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func (debugLogger) Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

func (debugLogger) Panicln(args ...interface{}) {
	logrus.Panicln(args...)
}
