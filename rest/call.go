package rest

import (
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/vndcontext"
)

type CallDelegate[REQ any] func(*logger.Logger, vndcontext.VndContext, *REQ) (*Result, error)

func Call[REQ any](e vndcontext.VndContext, req *REQ, name string, delegate CallDelegate[REQ]) (*Result, error) {
	log := logger.GetLogger(name)
	defer func() {
		log.Infow("completed")
		log.Sync()
	}()

	requestId := e.RequestId()
	log.With([]interface{}{
		"request_id", requestId,
	}...)

	return delegate(log, e, req)
}
