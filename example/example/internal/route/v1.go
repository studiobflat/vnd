package route

import (
	"example/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thienhaole92/vnd/mdw"
	"github.com/thienhaole92/vnd/rest"
	"github.com/thienhaole92/vnd/runner"
)

type V1 struct {
	*echo.Group
	auth mdw.AuthProvider
}

func (v1 *V1) Configure(rn *runner.Runner) error {
	fb, err := rn.GetInfra().Firebase()
	if err != nil {
		return err
	}
	v1.auth = fb.Auth

	s := service.NewService()
	return v1.registerRoutes(s)
}

func (v1 *V1) registerRoutes(s *service.Service) error {
	v1.GET("/", rest.Wrapper(s.GetSuccess), mdw.FirebaseAuth(middleware.DefaultSkipper, v1.auth))

	return nil
}
