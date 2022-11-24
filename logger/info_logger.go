package logger

import (
	"github.com/RichardKnop/logging"
	"github.com/sirupsen/logrus"
)

type infoLogger struct{}

func NewInfoLogger() logging.LoggerInterface {
	return &infoLogger{}
}

func (infoLogger) Print(args ...interface{}) {
	logrus.Info(args...)
}

func (infoLogger) Printf(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func (infoLogger) Println(args ...interface{}) {
	logrus.Infoln(args...)
}

func (infoLogger) Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func (infoLogger) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func (infoLogger) Fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}

func (infoLogger) Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func (infoLogger) Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

func (infoLogger) Panicln(args ...interface{}) {
	logrus.Panicln(args...)
}
