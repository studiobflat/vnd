package rest

import (
	gocontext "context"

	"github.com/labstack/echo/v4"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/mdw"
)

type Context struct {
	gocontext.Context
}

type CallDelegate[REQ any] func(*logger.Logger, *Context, *REQ) (*Result, error)

func Call[REQ any](e echo.Context, req *REQ, name string, delegate CallDelegate[REQ]) (*Result, error) {
	log := logger.GetLogger(name)
	defer func() {
		log.Infow("complete")
		log.Sync()
	}()

	ctx := e.Request().Context()
	requestId := e.Get(mdw.RequestIDContextKey)
	log.With([]interface{}{
		"request_id", requestId,
	}...)

	return delegate(log, &Context{
		Context: ctx,
	}, req)
}
