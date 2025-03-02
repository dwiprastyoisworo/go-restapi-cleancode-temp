package main

import (
	"fmt"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/internal/routes"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/configs"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/helpers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func AppInit() *http.Server {
	// setup logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	// setup user config
	cfg, err := configs.UserConfigInit()
	if err != nil {
		panic(err)
	}

	// setup postgres connection
	db, err := configs.PostgresInit(cfg.Postgres)
	if err != nil {
		panic(err)
	}

	// setup i18n
	bundle := configs.InitI18n()

	// set log level
	gin.SetMode(cfg.App.LogLevel)

	// create new gin server
	app := gin.New()

	// setup routes with custom logger
	app.Use(helpers.CustomLogger(logger))
	app.Use(helpers.SecurityHeaders())
	app.Use(helpers.CorsConfig())

	// setup routes
	appRoute := routes.NewUserRoute(db, bundle, app, cfg)
	appRoute.RouteInit()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.Port),
		Handler: app.Handler(),
	}
	return srv
}
