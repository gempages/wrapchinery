package logger

import (
	"github.com/RichardKnop/logging"
	"github.com/sirupsen/logrus"
)

type errorLogger struct{}

func NewErrorLogger() logging.LoggerInterface {
	return &errorLogger{}
}

func (errorLogger) Print(i ...interface{}) {
	logrus.Error(i)
}

func (errorLogger) Printf(s string, i ...interface{}) {
	logrus.Errorf(s, i)
}

func (errorLogger) Println(i ...interface{}) {
	logrus.Errorln(i)
}

func (errorLogger) Fatal(i ...interface{}) {
	logrus.Fatal(i)
}

func (errorLogger) Fatalf(s string, i ...interface{}) {
	logrus.Fatalf(s, i)
}

func (errorLogger) Fatalln(i ...interface{}) {
	logrus.Fatalln(i)
}

func (errorLogger) Panic(i ...interface{}) {
	logrus.Panic(i)
}

func (errorLogger) Panicf(s string, i ...interface{}) {
	logrus.Panicf(s, i)
}

func (errorLogger) Panicln(i ...interface{}) {
	logrus.Panicln(i)
}
