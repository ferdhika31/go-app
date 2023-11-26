package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ferdhika31/go-app/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type App struct {
	name         string
	port         string
	routeService string
	e            *echo.Echo
	server       *http.Server
	database     *sql.DB
}

// NewApp creates a new instance of the App struct with the given appName and appPort.
// It returns a pointer to the created App struct.
func NewApp(appName string, appPort string) *App {
	return &App{
		name: appName,
		port: appPort,
		e:    echo.New(),
		server: &http.Server{
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			IdleTimeout:  5 * time.Second,
		},
	}
}

// SetTimeZone sets the timezone of the application.
// It takes a string parameter `timezone` which is the timezone to be set.
// It loads the location of the timezone and sets the local time to the loaded location.
func (a *App) SetTimeZone(timezone string) {
	loc, _ := time.LoadLocation(timezone)
	time.Local = loc
}

// Run starts the application server and listens for incoming requests.
// It initializes the logger, sets the port, and runs the service.
// It returns an error if the server fails to start or encounters an error.
func (a *App) Run() {
	// Init logger
	logger.InitLogger()
	// Set port
	listenerPort := fmt.Sprintf(":%s", a.port)
	// Set default route
	a.setupRoute()
	// Run service
	a.server.Addr = listenerPort
	a.server.Handler = a.e
	logger.Fatalf("%v", a.e.Start(listenerPort).Error())
}

// ShutDown gracefully shuts down the server and closes the database connection.
// It sets a 30-second timeout for the server to finish the request it is currently handling.
// If the server fails to shut down gracefully, it logs an error message.
func (a *App) ShutDown() {
	logger.Infof("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	a.server.SetKeepAlivesEnabled(false)

	if err := a.server.Shutdown(ctx); err != nil {
		logger.Infof("%v", fmt.Sprintf("Could not gracefully shutdown the server %s: %v\n", a.name, err))
	}

	// closed database connection
	logger.Infof("Close connection database.")
	a.database.Close()
}

func (a *App) SetDatabase(db *sql.DB) {
	a.database = db
}
