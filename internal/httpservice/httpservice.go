package httpservice

import (
	"github.com/dliakhov/db-query-analyzer/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
		ErrorHandler: ErrorHandler,
	})

	appConfig := newApplicationConfig(conf, db)

	app.Use(logger.New())

	initRoutes(app, appConfig)

	return app.Listen(addr)
}

func initRoutes(app *fiber.App, appConfig *applicationConfig) {
	app.Get("/health", healthEndpoint(appConfig.db))

	app.Get("/v1/query", appConfig.QueryAnalyzerHandler.GetQueries)
}

func healthEndpoint(db *sqlx.DB) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		err := db.Ping()
		if err != nil {
			return err
		}
		ctx.Status(http.StatusOK)
		return ctx.SendString("OK")
	}
}
