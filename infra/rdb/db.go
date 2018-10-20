package rdb

import "github.com/gobuffalo/pop"

func NewConnection(connName string) (*pop.Connection, error) {
	return pop.Connect(connName)
}
