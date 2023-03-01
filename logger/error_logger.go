package logger

import (
	"github.com/RichardKnop/logging"
	"github.com/sirupsen/logrus"
)

type errorLogger struct{}

func NewErrorLogger() logging.LoggerInterface {
	return &errorLogger{}
}

func (errorLogger) Print(args ...interface{}) {
	logrus.Error(args...)
}

func (errorLogger) Printf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func (errorLogger) Println(args ...interface{}) {
	logrus.Errorln(args...)
}

func (errorLogger) Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func (errorLogger) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func (errorLogger) Fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}

func (errorLogger) Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func (errorLogger) Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

func (errorLogger) Panicln(args ...interface{}) {
	logrus.Panicln(args...)
}
