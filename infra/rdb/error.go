package rdb

import "errors"

var (
	ErrNotFound = errors.New("mysql select one: sql: no rows in result set")
)
