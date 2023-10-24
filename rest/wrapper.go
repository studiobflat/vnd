package rest

import (
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/labstack/echo/v4"
	vnderror "github.com/thienhaole92/vnd/error"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/mdw"
)

const RequestObjectContextKey = "service_requestObject"

func Wrapper[TREQ any](wrapped func(echo.Context, *TREQ) (*Result, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.GetLogger("Wrapper")
		defer log.Sync()

		requestId := c.Get(mdw.RequestIDContextKey)
		handler := runtime.FuncForPC(reflect.ValueOf(wrapped).Pointer()).Name()
		log.Infow("request begin",
			"request_id", requestId,
			"at", time.Now().Format(time.RFC3339),
			"path", c.Request().RequestURI,
			"handler", handler,
		)

		var req TREQ
		if err := c.Bind(&req); err != nil {
			log.Errorw("fail to bind request", "request_uri", c.Request().RequestURI, "err", err)
			return &vnderror.Error{CustomCode: -40001, HTTPError: &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid request"}}
		}

		if err := c.Validate(&req); err != nil {
			log.Errorw("fail to validate request", "request_uri", c.Request().RequestURI, "request_object", req, "err", err)
			return &vnderror.Error{CustomCode: -40002, HTTPError: &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid request"}}
		}

		c.Set(RequestObjectContextKey, req)

		res, err := wrapped(c, &req)
		if err != nil {
			return err
		}

		status := c.Response().Status
		if status != 0 {
			log.Infow("request end", "request_id", requestId, "at", time.Now().Format(time.RFC3339), "status", status)

			return c.JSON(
				status,
				Result{
					Data:       &res.Data,
					Pagination: res.Pagination,
				},
			)
		}

		log.Infow("request end", "request_id", requestId, "at", time.Now().Format(time.RFC3339), "status", http.StatusOK)

		return c.JSON(http.StatusOK, res)
	}
}
