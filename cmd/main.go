package main

import (
	"fmt"
	"github.com/ferdhika31/go-app/config"
	"github.com/ferdhika31/go-app/pkg/app"
	"github.com/ferdhika31/go-app/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load configuration
	configuration, _ := config.LoadConfig()
	appName := configuration.APP.Name
	appPort := configuration.APP.Port
	appRouteService := configuration.APP.RouteService

	// Init app
	newApp := app.NewApp(appName, appPort)

	// Init database
	config.InitDatabasePostgres(newApp)

	// Set route service
	newApp.SetRouteService(appRouteService)

	// Run newApp
	newApp.Run()

	// Wait for interrupt signal to gracefully shut down the server
	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	newApp.ShutDown()
	logger.Infof(fmt.Sprintf("Server %s stopped", appName), "gracefully", "")
}
