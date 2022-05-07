package httpservice

import (
	"github.com/dliakhov/db-query-analyzer/config"
	"github.com/dliakhov/db-query-analyzer/internal/repository"
	"github.com/jmoiron/sqlx"
)

type applicationConfig struct {
	config               *config.Config
	QueryAnalyzerHandler *QueryAnalyzerHandler
}

func newApplicationConfig(config *config.Config, db *sqlx.DB) *applicationConfig {
	return &applicationConfig{
		config:               config,
		QueryAnalyzerHandler: NewQueryAnalyzer(repository.NewQueryAnalyzerRepository(db)),
	}
}
