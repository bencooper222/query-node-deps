package db

import (
	"github.com/bencooper222/query-node-deps/pkg/env"
	"github.com/bencooper222/query-node-deps/pkg/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BuildPostgresClient(host string, user string, password string, dbname string, port string) (gorm.DB, error) {
	dsn := "host=" + host + " user=" + user + " dbname=" + dbname + " port=" + port + " password=" + password
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return *db, err
}

func BuildPostgresClientFromEnvOrDie() gorm.DB {
	db, err := BuildPostgresClient(env.Get("POSTGRES_HOST", "localhost"), env.Get("POSTGRES_USER", "postgres"), env.Get("POSTGRES_PASSWORD", "password"), env.Get("POSTGRES_DBNAME", "postgres"), env.Get("POSTGRES_PORT", "5432"))
	util.CheckErr(err)

	return db
}
