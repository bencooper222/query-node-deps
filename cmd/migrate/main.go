package main

import (
	"log"

	db "github.com/bencooper222/query-node-deps/pkg/db"
	dbTypes "github.com/bencooper222/query-node-deps/pkg/db/types"
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

		{
			ID: "1658718107",
			Migrate: func(tx *gorm.DB) error {

				return tx.AutoMigrate(&dbTypes.Repository{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("repositories")
			},
		},

		{
			ID: "1658719633",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&dbTypes.Commit{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("commits")
			},
		},

		{
			ID: "1658722299",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&dbTypes.Dependency{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("dependencies")
			},
		},
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Finished DB migration")
}
