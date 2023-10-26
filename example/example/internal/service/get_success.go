package service

import (
	_ "github.com/thienhaole92/vnd/error"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/rest"
	"github.com/thienhaole92/vnd/vndcontext"
)

type GetSuccessReq struct {
}

func (s *Service) GetSuccess(e vndcontext.VndContext, req *GetSuccessReq) (*rest.Result, error) {
	delegate := func(log *logger.Logger, ctx vndcontext.VndContext, req *GetSuccessReq) (*rest.Result, error) {
		exec := NewGetSuccess(log)
		return exec.Execute(ctx, req)
	}
	return rest.Call[GetSuccessReq](e, req, "GetSuccess", delegate)
}

type getSuccess struct {
	log *logger.Logger
}

func NewGetSuccess(log *logger.Logger) *getSuccess {
	return &getSuccess{
		log: log,
	}
}

func (s *getSuccess) Execute(ctx vndcontext.VndContext, req *GetSuccessReq) (*rest.Result, error) {
	return &rest.Result{}, nil
}
