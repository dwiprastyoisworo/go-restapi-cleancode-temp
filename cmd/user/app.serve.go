package main

import (
	"context"
	"fmt"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/configs"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/helpers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	// setup user config
	cfg, err := configs.UserConfigInit()
	if err != nil {
		panic(err)
	}

	// set log level
	gin.SetMode(cfg.App.LogLevel)

	// create new gin server
	app := gin.New()

	// setup routes with custom logger
	app.Use(helpers.CustomLogger(logger))
	app.Use(helpers.SecurityHeaders())
	app.Use(helpers.CorsConfig())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.Port),
		Handler: app.Handler(),
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		// service connections
		log.Println("Starting Server ...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")

}
