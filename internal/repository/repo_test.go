package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dliakhov/db-query-analyzer/internal/models"
	"github.com/jmoiron/sqlx"
	"reflect"
	"testing"
	"time"
)

func TestQueryAnalyzerRepositoryImpl_GetDatabaseQueryInfo(t *testing.T) {
	rows := []string{"id", "created_at", "updated_at", "deleted_at", "query", "execution_time_ms"}

	const layout = "2006-01-02 15:04:05"
	queryTime, _ := time.Parse(layout, "2020-01-01 13:00:00")
	type fields struct {
		newMockDB func() (*sqlx.DB, sqlmock.Sqlmock)
	}
	type args struct {
		queryRequest models.QueryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.DatabaseQueryInfo
		wantErr bool
	}{
		{
			name: "should return query info",
			fields: fields{
				newMockDB: func() (*sqlx.DB, sqlmock.Sqlmock) {
					db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}
					mock.ExpectQuery(`SELECT * FROM query_analyzer.database_query_info WHERE deleted_at IS NULL ORDER BY execution_time_ms ASC LIMIT $1 OFFSET $2`).
						WithArgs(5, 0).
						WillReturnRows(sqlmock.NewRows(rows).AddRow(
							1, queryTime, queryTime, nil, "SELECT * FROM table", 100,
						))
					return sqlx.NewDb(db, "postgres"), mock
				},
			},
			args: args{
				queryRequest: models.QueryRequest{
					Page:      1,
					Size:      5,
					QueryType: "",
					OrderBy:   "",
				},
			},
			want: []models.DatabaseQueryInfo{
				{
					Model: models.Model{
						ID:        1,
						CreatedAt: queryTime,
						UpdatedAt: queryTime,
						DeletedAt: nil,
					},
					Query:           "SELECT * FROM table",
					ExecutionTimeMs: 100,
				},
			},
		},
		{
			name: "should return query info and filter by query type",
			fields: fields{
				newMockDB: func() (*sqlx.DB, sqlmock.Sqlmock) {
					db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}
					mock.ExpectQuery(`SELECT * FROM query_analyzer.database_query_info 
							WHERE (query LIKE 'insert%' OR query LIKE 'INSERT%') AND deleted_at IS NULL 
							ORDER BY execution_time_ms ASC 
							LIMIT $1 OFFSET $2`).
						WithArgs(10, 10).
						WillReturnRows(sqlmock.NewRows(rows).AddRow(
							1, queryTime, queryTime, nil, "INSERT INTO users (id, name, age) VALUES (?, ?, ?)", 200,
						))
					return sqlx.NewDb(db, "postgres"), mock
				},
			},
			args: args{
				queryRequest: models.QueryRequest{
					Page:      2,
					Size:      10,
					QueryType: "insert",
					OrderBy:   "",
				},
			},
			want: []models.DatabaseQueryInfo{
				{
					Model: models.Model{
						ID:        1,
						CreatedAt: queryTime,
						UpdatedAt: queryTime,
						DeletedAt: nil,
					},
					Query:           "INSERT INTO users (id, name, age) VALUES (?, ?, ?)",
					ExecutionTimeMs: 200,
				},
			},
		},
		{
			name: "should return query info and filter by query type and order by execution_time_ms descending",
			fields: fields{
				newMockDB: func() (*sqlx.DB, sqlmock.Sqlmock) {
					db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}
					mock.ExpectQuery(`SELECT * FROM query_analyzer.database_query_info 
							WHERE (query LIKE 'insert%' OR query LIKE 'INSERT%') AND deleted_at IS NULL 
							ORDER BY execution_time_ms DESC 
							LIMIT $1 OFFSET $2`).
						WithArgs(10, 10).
						WillReturnRows(sqlmock.NewRows(rows).AddRow(
							1, queryTime, queryTime, nil, "INSERT INTO users (id, name, age) VALUES (?, ?, ?)", 200,
						))
					return sqlx.NewDb(db, "postgres"), mock
				},
			},
			args: args{
				queryRequest: models.QueryRequest{
					Page:      2,
					Size:      10,
					QueryType: "insert",
					OrderBy:   "desc",
				},
			},
			want: []models.DatabaseQueryInfo{
				{
					Model: models.Model{
						ID:        1,
						CreatedAt: queryTime,
						UpdatedAt: queryTime,
						DeletedAt: nil,
					},
					Query:           "INSERT INTO users (id, name, age) VALUES (?, ?, ?)",
					ExecutionTimeMs: 200,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := tt.fields.newMockDB()
			r := NewQueryAnalyzerRepository(db)
			got, err := r.GetDatabaseQueryInfo(tt.args.queryRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDatabaseQueryInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDatabaseQueryInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
