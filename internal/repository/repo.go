package repository

import (
	"github.com/dliakhov/db-query-analyzer/internal/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

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

	likeQuery := strings.ToUpper(queryRequest.QueryType) + "%"
	selectQuery := `SELECT * FROM query_analyzer.database_query_info 
         WHERE query like $1 AND deleted_at IS NULL`
	if strings.ToLower(queryRequest.OrderBy) == "asc" {
		selectQuery += ` ORDER BY execution_time_ms ASC`
	} else {
		selectQuery += ` ORDER BY execution_time_ms DESC`
	}
	selectQuery += ` LIMIT $2 OFFSET $3`

	err := r.db.Select(&queries, selectQuery, likeQuery, queryRequest.Size, (queryRequest.Page-1)*queryRequest.Size)
	if err != nil {
		return nil, err
	}
	return queries, nil

}
