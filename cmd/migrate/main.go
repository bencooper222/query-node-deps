package main

import (
	"log"

	"github.com/bencooper222/query-node-deps/pkg/db"
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
)

func main() {
	pgclient, err := db.BuildPostgresClient("localhost", "postgres", "password", "postgres", "5432")

	if err != nil {
		panic(err)
	}
	log.Println("Starting DB migration")
	m := gormigrate.New(&pgclient, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "1658713212",
			Migrate: func(tx *gorm.DB) error {
				rawDb, _ := tx.DB()
				_, err = rawDb.Exec("CREATE EXTENSION IF NOT EXISTS semver")
				return err
			},
			Rollback: func(tx *gorm.DB) error {
				rawDb, _ := tx.DB()
				_, err = rawDb.Exec("DELETE EXTENSION IF EXIST semver")
				return err
			},
		},
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Finished DB migration")
}
