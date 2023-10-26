package route

import (
	"example-rest/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/thienhaole92/vnd/rest"
	"github.com/thienhaole92/vnd/runner"
)

type V1 struct {
	*echo.Group
}

func (v1 *V1) Configure(rn *runner.Runner) error {
	s := service.NewService()

	return v1.registerRoutes(s)
}

func (v1 *V1) registerRoutes(s *service.Service) error {
	v1.GET("/get-success/:id", rest.Wrapper(s.GetSuccess))
	v1.GET("/get-fail", rest.Wrapper(s.GetFail))

	return nil
}
