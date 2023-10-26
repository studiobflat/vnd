package event

import (
	"context"

	"github.com/thienhaole92/vnd/logger"
)

type ConsumeDelegate[D any] func(*logger.Logger, context.Context, *EventSchema[D]) error

func Consume[D any](ctx context.Context, message EventString, name string, delegate ConsumeDelegate[D]) error {
	log := logger.GetLogger(name)
	defer func() {
		log.Infow("completed")
		log.Sync()
	}()

	schema := EventSchema[D]{}
	if err := message.UnpackEvent(&schema); err != nil {
		return err
	}

	return delegate(log, ctx, &schema)
}
