package event

import (
	gocontext "context"

	"github.com/thienhaole92/vnd/logger"
)

type Context struct {
	gocontext.Context
}

type ConsumeDelegate[D any] func(*logger.Logger, *Context, *EventSchema[D]) error

func Consume[D any](ctx gocontext.Context, message EventString, name string, delegate ConsumeDelegate[D]) error {
	log := logger.GetLogger(name)
	defer func() {
		log.Infow("complete")
		log.Sync()
	}()

	schema := EventSchema[D]{}
	if err := message.UnpackEvent(&schema); err != nil {
		return err
	}

	return delegate(log, &Context{
		Context: ctx,
	}, &schema)
}
