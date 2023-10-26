package service

import (
	"context"

	"github.com/labstack/echo/v4"
	_ "github.com/thienhaole92/vnd/error"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/rest"
)

type GetSuccessReq struct {
	Id string `param:"id" validate:"required"`
}

// Get success godoc
//
//	@Summary		Get success
//	@Description	Gets success description.
//	@Tags			get
//	@Consume		json
//	@Produce		json
//	@param 			Authorization 	header 		string	true 	"Your access token" default(Bearer <Add access token here>)
//	@Param			id				path		string	true	"ID"
//	@Success		200				{object}	rest.DocResult[string]	"Return data"
//	@Failure		400				{object}	error.Error	"Error object"
//	@Router			/get-success/{id} [get]
//	@Security		Bearer
func (s *Service) GetSuccess(e echo.Context, req *GetSuccessReq) (*rest.Result, error) {
	delegate := func(log *logger.Logger, ctx *rest.Context, req *GetSuccessReq) (*rest.Result, error) {
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

func (s *getSuccess) Execute(ctx context.Context, req *GetSuccessReq) (*rest.Result, error) {
	return &rest.Result{
		Data: req.Id,
	}, nil
}
