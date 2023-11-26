package app

import (
	"fmt"
	response "github.com/ferdhika31/go-app/pkg/response"
	"github.com/labstack/echo/v4"
)

type Router struct {
	router *echo.Group
}

func (a *App) SetRouteService(routeService string) {
	a.routeService = routeService
}

func (a *App) setupRoute() *Router {
	// Set echo instance
	e := a.e
	// Set route not found handler
	e.RouteNotFound(fmt.Sprintf("/*"), func(e echo.Context) error {
		return response.Json(e, response.RouteNotFound, nil)
	})
	// Get prefix
	prefix := a.routeService
	// Set init prefix route
	appName := a.name
	publicRoute := e.Group(prefix)
	publicRoute.GET("/health-check", func(e echo.Context) error {
		return response.JsonWithMessage(e, response.DataFetched, nil, "Service "+appName+" is running")
	})
	//
	return &Router{
		router: publicRoute,
	}
}
