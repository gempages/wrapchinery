package logger

import (
	"context"
	"fmt"

	"github.com/RichardKnop/logging"
	"github.com/gempages/go-helper/log"
	"github.com/sirupsen/logrus"
)

type fatalLogger struct{}

func NewFatalLogger() logging.LoggerInterface {
	return &fatalLogger{}
}

func (fatalLogger) Print(args ...interface{}) {
	err := toError(args...)
	if err != nil {
		log.LogError(context.Background(), err)
	}
	logrus.Print(args...)
}

func (fatalLogger) Printf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	log.LogError(context.Background(), err)
	logrus.Printf(format, args...)
}

func (fatalLogger) Println(args ...interface{}) {
	err := toError(args...)
	if err != nil {
		log.LogError(context.Background(), err)
	}
	logrus.Println(args...)
}

func (fatalLogger) Fatal(args ...interface{}) {
	err := toError(args...)
	if err != nil {
		log.LogError(context.Background(), err)
	}
	logrus.Fatal(args...)
}

func (fatalLogger) Fatalf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	log.LogError(context.Background(), err)
	logrus.Fatalf(format, args...)
}

func (fatalLogger) Fatalln(args ...interface{}) {
	err := toError(args...)
	if err != nil {
		log.LogError(context.Background(), err)
	}
	logrus.Fatalln(args...)
}

func (fatalLogger) Panic(args ...interface{}) {
	err := toError(args...)
	if err != nil {
		log.LogError(context.Background(), err)
	}
	logrus.Panic(args...)
}

func (fatalLogger) Panicf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	log.LogError(context.Background(), err)
	logrus.Panicf(format, args...)
}

func (fatalLogger) Panicln(args ...interface{}) {
	err := toError(args...)
	if err != nil {
		log.LogError(context.Background(), err)
	}
	logrus.Panicln(args...)
}
