package wrapchinery

import (
	"context"
	"fmt"

	"github.com/RichardKnop/machinery/v2"
	"github.com/gempages/go-helper/log"
)

func NewWorker(concurrency int) *machinery.Worker {
	w := server.WrapNewWorker(concurrency)
	// Override default error handler function in order to remove task ID in error message,
	// which makes it easier to manage issues in Sentry
	w.SetErrorHandler(func(err error) {
		log.Error(context.Background(), fmt.Errorf("worker task failed: %w", err))
	})
	return w
}
