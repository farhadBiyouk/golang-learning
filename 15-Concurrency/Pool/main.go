package main

import (
	"fmt"
	"sync"
)

type DbConnection struct {
	Host     string
	Name     string
	User     string
	Password string
}

var ConnectionPool sync.Pool = sync.Pool{
	New: func() interface{} {
		return &DbConnection{
			Host:     "localhost",
			Name:     "testdb",
			User:     "root",
			Password: "root",
		}
	},
}

func main() {
	connection := ConnectionPool.Get().(*DbConnection)
	fmt.Println(connection)

	ConnectionPool.Put(connection)
}
