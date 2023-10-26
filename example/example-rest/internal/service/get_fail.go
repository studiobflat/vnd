package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	vnderror "github.com/thienhaole92/vnd/error"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/rest"
)

type GetFailReq struct {
}

func (s *Service) GetFail(e echo.Context, req *GetFailReq) (*rest.Result, error) {
	delegate := func(log *logger.Logger, ctx *rest.Context, req *GetFailReq) (*rest.Result, error) {
		exec := NewGetFail(log)
		return exec.Execute(ctx, req)
	}
	return rest.Call[GetFailReq](e, req, "GetFail", delegate)
}

type getFail struct {
	log *logger.Logger
}

func NewGetFail(log *logger.Logger) *getFail {
	return &getFail{
		log: log,
	}
}

func (s *getFail) Execute(ctx context.Context, req *GetFailReq) (*rest.Result, error) {
	return nil, &vnderror.Error{
		CustomCode: -50011,
		HTTPError: &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  "something wrong",
			Internal: fmt.Errorf("something wrong internal"),
		},
	}
}
