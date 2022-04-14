package logger

import (
	"github.com/RichardKnop/logging"
	"github.com/sirupsen/logrus"
)

type fatalLogger struct{}

func NewFatalLogger() logging.LoggerInterface {
	return &fatalLogger{}
}

func (fatalLogger) Print(i ...interface{}) {
	logrus.Print(i)
}

func (fatalLogger) Printf(s string, i ...interface{}) {
	logrus.Printf(s, i)
}

func (fatalLogger) Println(i ...interface{}) {
	logrus.Println(i)
}

func (fatalLogger) Fatal(i ...interface{}) {
	logrus.Fatal(i)
}

func (fatalLogger) Fatalf(s string, i ...interface{}) {
	logrus.Fatalf(s, i)
}

func (fatalLogger) Fatalln(i ...interface{}) {
	logrus.Fatalln(i)
}

func (fatalLogger) Panic(i ...interface{}) {
	logrus.Panic(i)
}

func (fatalLogger) Panicf(s string, i ...interface{}) {
	logrus.Panicf(s, i)
}

func (fatalLogger) Panicln(i ...interface{}) {
	logrus.Panicln(i)
}
