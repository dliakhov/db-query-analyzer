package httpservice

import (
	"github.com/dliakhov/db-query-analyzer/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net"
	"net/http"
	"time"
)

func Run(conf *config.Config, db *sqlx.DB) error {
	addr := net.JoinHostPort(conf.HTTPServer.HostName, conf.HTTPServer.Port)
	app := fiber.New(fiber.Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 120 * time.Second,
	})

	appConfig := newApplicationConfig(conf, db)
	initRoutes(app, appConfig)

	return app.Listen(addr)
}

func initRoutes(app *fiber.App, appConfig *applicationConfig) {
	app.Get("/health", healthEndpoint())

	app.Get("/v1/query", appConfig.QueryAnalyzerHandler.GetQueries)
}

func healthEndpoint() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Status(http.StatusOK)
		return ctx.SendString("OK")
	}
}
