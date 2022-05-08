package repository

import (
	"fmt"
	"github.com/dliakhov/db-query-analyzer/internal/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

//go:generate mockgen -source=repo.go -destination=repo_mock.go -package=repository
type QueryAnalyzerRepository interface {
	GetDatabaseQueryInfo(queryRequest models.QueryRequest) ([]models.DatabaseQueryInfo, error)
}

type QueryAnalyzerRepositoryImpl struct {
	db *sqlx.DB
}

func NewQueryAnalyzerRepository(db *sqlx.DB) *QueryAnalyzerRepositoryImpl {
	return &QueryAnalyzerRepositoryImpl{db: db}
}

func (r *QueryAnalyzerRepositoryImpl) GetDatabaseQueryInfo(queryRequest models.QueryRequest) ([]models.DatabaseQueryInfo, error) {
	queries := make([]models.DatabaseQueryInfo, 0)

	selectQuery := `SELECT * FROM query_analyzer.database_query_info WHERE `
	if queryRequest.QueryType != "" {
		likeQuery := fmt.Sprintf(`(query LIKE '%s%%' OR query LIKE '%s%%') AND`, strings.ToLower(queryRequest.QueryType), strings.ToUpper(queryRequest.QueryType))
		selectQuery += likeQuery
	}
	selectQuery += ` deleted_at IS NULL`

	if strings.ToLower(queryRequest.ExecutionTimeSort) == "desc" {
		selectQuery += ` ORDER BY execution_time_ms DESC`
	} else {
		selectQuery += ` ORDER BY execution_time_ms ASC`
	}
	selectQuery += ` LIMIT $1 OFFSET $2`

	err := r.db.Select(&queries, selectQuery, queryRequest.Size, (queryRequest.Page-1)*queryRequest.Size)
	if err != nil {
		return nil, err
	}
	return queries, nil

}
