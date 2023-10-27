package main

import (
	"example/internal/route"

	"github.com/labstack/echo/v4"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/postgres"
	"github.com/thienhaole92/vnd/runner"
)

func main() {
	options := []runner.RunnerOption{
		runner.NewInfraHookOption("postgres_db", runner.PostgresHook),
		runner.NewInfraHookOption("firebase", runner.FirebaseHook),
		runner.BuildPostgresDatabaseMigrationHook(migrateDatabaseHook),
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

func migrateDatabaseHook() (string, string, error) {
	c, err := postgres.NewConfig()
	if err != nil {
		return "", "", err
	}

	source := "file://db/migrations"
	databaseURL := c.Url
	return source, databaseURL, nil
}
