package main

import (
	"github.com/bencooper222/query-node-deps/pkg/db"
)

func main() {
	pgclient, err := db.BuildPostgresClient("localhost", "postgres", "password", "postgres", "5432")

	if err != nil {
		panic(err)
	}

	type TestTable struct {
		uid int
	}
	pgclient.AutoMigrate(&TestTable{})

	pgclient.Table("test_tables").Create(&TestTable{uid: 1})

}
