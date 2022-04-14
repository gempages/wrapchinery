package logger

import (
	"github.com/RichardKnop/logging"
	"github.com/sirupsen/logrus"
)

type infoLogger struct{}

func NewInfoLogger() logging.LoggerInterface {
	return &infoLogger{}
}

func (infoLogger) Print(i ...interface{}) {
	logrus.Info(i)
}

func (infoLogger) Printf(s string, i ...interface{}) {
	logrus.Infof(s, i)
}

func (infoLogger) Println(i ...interface{}) {
	logrus.Infoln(i)
}

func (infoLogger) Fatal(i ...interface{}) {
	logrus.Fatal(i)
}

func (infoLogger) Fatalf(s string, i ...interface{}) {
	logrus.Fatalf(s, i)
}

func (infoLogger) Fatalln(i ...interface{}) {
	logrus.Fatalln(i)
}

func (infoLogger) Panic(i ...interface{}) {
	logrus.Panic(i)
}

func (infoLogger) Panicf(s string, i ...interface{}) {
	logrus.Panicf(s, i)
}

func (infoLogger) Panicln(i ...interface{}) {
	logrus.Panicln(i)
}
