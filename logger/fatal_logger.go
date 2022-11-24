package logger

import (
	"github.com/RichardKnop/logging"
	"github.com/sirupsen/logrus"
)

type fatalLogger struct{}

func NewFatalLogger() logging.LoggerInterface {
	return &fatalLogger{}
}

func (fatalLogger) Print(args ...interface{}) {
	logrus.Print(args...)
}

func (fatalLogger) Printf(format string, args ...interface{}) {
	logrus.Printf(format, args...)
}

func (fatalLogger) Println(args ...interface{}) {
	logrus.Println(args...)
}

func (fatalLogger) Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func (fatalLogger) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func (fatalLogger) Fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}

func (fatalLogger) Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func (fatalLogger) Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

func (fatalLogger) Panicln(args ...interface{}) {
	logrus.Panicln(args...)
}
