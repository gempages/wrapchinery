package logger

import (
	"github.com/RichardKnop/logging"
	"github.com/sirupsen/logrus"
)

type debugLogger struct{}

func NewDebugLogger() logging.LoggerInterface {
	return &debugLogger{}
}

func (debugLogger) Print(i ...interface{}) {
	logrus.Debug(i)
}

func (debugLogger) Printf(s string, i ...interface{}) {
	logrus.Debugf(s, i)
}

func (debugLogger) Println(i ...interface{}) {
	logrus.Debugln(i)
}

func (debugLogger) Fatal(i ...interface{}) {
	logrus.Fatal(i)
}

func (debugLogger) Fatalf(s string, i ...interface{}) {
	logrus.Fatalf(s, i)
}

func (debugLogger) Fatalln(i ...interface{}) {
	logrus.Fatalln(i)
}

func (debugLogger) Panic(i ...interface{}) {
	logrus.Panic(i)
}

func (debugLogger) Panicf(s string, i ...interface{}) {
	logrus.Panicf(s, i)
}

func (debugLogger) Panicln(i ...interface{}) {
	logrus.Panicln(i)
}
