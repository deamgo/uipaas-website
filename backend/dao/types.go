package dao

import "github.com/pkg/errors"

var (
	// DBError indicates a database error.
	DBError = errors.New("database error")
)
