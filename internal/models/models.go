package models

import "time"

type Model struct {
	ID        uint       `db:"id" json:"-"`
	CreatedAt time.Time  `db:"created_at" json:"-"`
	UpdatedAt time.Time  `db:"updated_at" json:"-"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
}

type DatabaseQueryInfo struct {
	Model
	Query           string `db:"query"`
	ExecutionTimeMs uint   `db:"execution_time_ms"`
}

type QueryRequest struct {
	Page      uint   `query:"page" validate:"required,min=1"`
	Size      uint   `query:"size" validate:"required,min=1"`
	QueryType string `query:"query_type" validate:"required,oneof=SELECT INSERT UPDATE DELETE select insert update delete"`
	OrderBy   string `query:"order_by" validate:"oneof=asc desc ASC DESC"`
}
