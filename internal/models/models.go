package models

import (
	"time"
)

type Model struct {
	ID        uint       `db:"id" json:"-"`
	CreatedAt time.Time  `db:"created_at" json:"-"`
	UpdatedAt time.Time  `db:"updated_at" json:"-"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
}

type DatabaseQueryInfo struct {
	Model
	Query           string `db:"query" json:"query"`
	ExecutionTimeMs uint   `db:"execution_time_ms" json:"execution_time_ms"`
}

type QueryRequest struct {
	Page              uint   `query:"page" validate:"required,min=1"`
	Size              uint   `query:"size" validate:"required,min=1"`
	QueryType         string `query:"query_type" validate:"omitempty,oneof=SELECT INSERT UPDATE DELETE select insert update delete"`
	ExecutionTimeSort string `query:"execution_time_sort" validate:"omitempty,oneof=asc desc ASC DESC"`
}
