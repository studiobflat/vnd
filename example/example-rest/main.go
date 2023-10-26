package main

import (
	"example-rest/internal/route"

	"github.com/labstack/echo/v4"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/runner"
)

// @title				Swagger Example API
// @version				1.0
// @description     	This is a sample server celler server.
// @contact.name		REST API developer
// @contact.email		developer@example.com
//
// @host      			localhost:7080
// @BasePath  			/v1
//
// @tag.name			get
// @tag.description		Get methods demo

func main() {
	options := []runner.RunnerOption{
		runner.BuildMonitorServerOption(runner.DefaultMonitorEchoHook),
		runner.BuildRestServerOption(restServiceHook),
	}
	r := runner.NewRunner(options...)
	r.Run()
}

func restServiceHook(rn *runner.Runner, e *echo.Echo, eg *echo.Group) error {
	log := logger.GetLogger("restServiceHook")
	defer log.Sync()

	v1 := &route.V1{Group: eg.Group("/v1")}
	err := v1.Configure(rn)
	if err != nil {
		log.Errorw("failed to get register routes", "error", err)
		return err
	}

	return nil
}
