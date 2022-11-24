package logger

import (
	"context"
	"fmt"

	"github.com/RichardKnop/logging"
	"github.com/gempages/go-helper/log"
	"github.com/sirupsen/logrus"
)

type errorLogger struct{}

func NewErrorLogger() logging.LoggerInterface {
	return &errorLogger{}
}

func (errorLogger) Print(args ...interface{}) {
	err := toError(args...)
	if err != nil {
		log.LogError(context.Background(), err)
	}
	logrus.Error(args...)
}

func (errorLogger) Printf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	log.LogError(context.Background(), err)
	logrus.Errorf(format, args...)
}

func (errorLogger) Println(args ...interface{}) {
	err := toError(args...)
	if err != nil {
		log.LogError(context.Background(), err)
	}
	logrus.Errorln(args...)
}

func (errorLogger) Fatal(args ...interface{}) {
	err := toError(args...)
	if err != nil {
		log.LogError(context.Background(), err)
	}
	logrus.Fatal(args...)
}

func (errorLogger) Fatalf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	log.LogError(context.Background(), err)
	logrus.Fatalf(format, args...)
}

func (errorLogger) Fatalln(args ...interface{}) {
	err := toError(args...)
	if err != nil {
		log.LogError(context.Background(), err)
	}
	logrus.Fatalln(args...)
}

func (errorLogger) Panic(args ...interface{}) {
	err := toError(args...)
	if err != nil {
		log.LogError(context.Background(), err)
	}
	logrus.Panic(args...)
}

func (errorLogger) Panicf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	log.LogError(context.Background(), err)
	logrus.Panicf(format, args...)
}

func (errorLogger) Panicln(args ...interface{}) {
	err := toError(args...)
	if err != nil {
		log.LogError(context.Background(), err)
	}
	logrus.Panicln(args...)
}

func toError(args ...interface{}) error {
	var err error
	for _, arg := range args {
		switch e := arg.(type) {
		case error:
			if err == nil {
				err = e
			} else {
				err = fmt.Errorf("%s %w", err.Error(), e)
			}
		default:
			if err == nil {
				err = fmt.Errorf("%v", e)
			} else {
				err = fmt.Errorf("%s %w", err.Error(), fmt.Errorf("%v", e))
			}
		}
	}
	return err
}
