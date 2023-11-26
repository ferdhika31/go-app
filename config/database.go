package config

import (
	"fmt"
	"github.com/ferdhika31/go-app/pkg/app"
	"github.com/ferdhika31/go-app/pkg/database"
)

func InitDatabasePostgres(a *app.App) {
	// get from env
	dbHost := GetEnvByKey("DB_HOST")
	dbPort := GetEnvByKey("DB_PORT")
	dbUser := GetEnvByKey("DB_USER")
	dbName := GetEnvByKey("DB_NAME")
	dbPassword := GetEnvByKey("DB_PASS")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password='%s'", dbHost, dbPort, dbUser, dbName, dbPassword)
	//ini connection
	dbPg := database.Postgresql(dsn, 10, 10)
	//set db apps
	a.SetDatabase(dbPg)
}
