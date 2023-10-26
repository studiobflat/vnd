package vndcontext

import (
	"context"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/thienhaole92/vnd/mdw"
)

type VndContext interface {
	RequestContext() context.Context
	RequestId() string
}

type VContext struct {
	echo.Context
}

func (c *VContext) RequestContext() context.Context {
	return c.Request().Context()
}

func (c *VContext) RequestId() string {
	id := c.Get(mdw.RequestIDContextKey)
	if id != nil && reflect.TypeOf(id).Name() == "string" {
		return id.(string)
	}

	return ""
}
