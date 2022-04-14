package logger

import (
	"github.com/RichardKnop/logging"
	"github.com/sirupsen/logrus"
)

type warningLogger struct{}

func NewWarningLogger() logging.LoggerInterface {
	return &warningLogger{}
}

func (warningLogger) Print(i ...interface{}) {
	logrus.Warning(i)
}

func (warningLogger) Printf(s string, i ...interface{}) {
	logrus.Warningf(s, i)
}

func (warningLogger) Println(i ...interface{}) {
	logrus.Warningln(i)
}

func (warningLogger) Fatal(i ...interface{}) {
	logrus.Fatal(i)
}

func (warningLogger) Fatalf(s string, i ...interface{}) {
	logrus.Fatalf(s, i)
}

func (warningLogger) Fatalln(i ...interface{}) {
	logrus.Fatalln(i)
}

func (warningLogger) Panic(i ...interface{}) {
	logrus.Panic(i)
}

func (warningLogger) Panicf(s string, i ...interface{}) {
	logrus.Panicf(s, i)
}

func (warningLogger) Panicln(i ...interface{}) {
	logrus.Panicln(i)
}
