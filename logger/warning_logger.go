package logger

import (
	"github.com/RichardKnop/logging"
	"github.com/sirupsen/logrus"
)

type warningLogger struct{}

func NewWarningLogger() logging.LoggerInterface {
	return &warningLogger{}
}

func (warningLogger) Print(args ...interface{}) {
	logrus.Warning(args...)
}

func (warningLogger) Printf(format string, args ...interface{}) {
	logrus.Warningf(format, args...)
}

func (warningLogger) Println(args ...interface{}) {
	logrus.Warningln(args...)
}

func (warningLogger) Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func (warningLogger) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func (warningLogger) Fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}

func (warningLogger) Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func (warningLogger) Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

func (warningLogger) Panicln(args ...interface{}) {
	logrus.Panicln(args...)
}
