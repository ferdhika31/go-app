package database

import (
	"database/sql"
	"github.com/ferdhika31/go-app/pkg/logger"
	"time"

	"github.com/lib/pq"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

var db *sql.DB

func Postgresql(dsn string, SetMaxIdleConns, SetMaxOpenConns int) *sql.DB {
	appName := "App Name"
	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithServiceName(appName))
	sqlDB, err := sqltrace.Open("postgres", dsn)
	if err != nil {
		logger.Panicf("Error connecting to database: %s", err.Error())
	}
	if err := sqlDB.Ping(); err != nil {
		logger.Panicf("Error connecting to database: %s", err.Error())
	}

	sqlDB.SetMaxIdleConns(SetMaxIdleConns)
	sqlDB.SetMaxOpenConns(SetMaxOpenConns)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	db = sqlDB

	return sqlDB
}

func GetPostgresDB() *sql.DB {
	return db
}
