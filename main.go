package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	conf "github.com/sammy1881/tinyurl/config"
	co "github.com/sammy1881/tinyurl/controller"
)

var (
	router *gin.Engine
)

func main() {

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.Use(setVariables)
	router.GET("/", co.Home)
	router.POST("/addurl", co.AddURL)
	router.GET("/:shorturl", co.GetURL)

	srv := &http.Server{
		Addr:    ":" + conf.GetConfig().Port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")

}

func setVariables(c *gin.Context) {
	c.Set("ShortenerHostname", conf.GetConfig().ShortenerHostname)
	c.Set("IdLength", conf.GetConfig().IdLength)
	c.Set("IdAlphabet", conf.GetConfig().IdAlphabet)
	c.Set("Port", conf.GetConfig().Port)
	c.Set("DB", conf.GetConfig().DB)
	c.Set("Bucket", conf.GetConfig().Bucket)
	c.Next()
}
