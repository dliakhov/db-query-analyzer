package models

import "time"

type Model struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type DatabaseQueryInfo struct {
	Model
	Query           string
	ExecutionTimeMs uint
}
