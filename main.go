package main

import (
	"binary-tree-max-path-sum/internal/tree"
	"binary-tree-max-path-sum/pkg/config"
	"binary-tree-max-path-sum/pkg/graceful"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Calculate max path sum service started")
	Execute()
}

func Execute() {

	// Load configuration depending on app environment
	configFile := fmt.Sprintf("./pkg/config/%s", "local")
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Loadconfig failed, %v", err)
	}

	// Set gin router
	router := gin.Default()
	tree.NewTreeHandler(router.Group(cfg.ServerConfig.RoutePrefix))

	// Create server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ServerConfig.Port),
		ReadTimeout:  time.Duration(int64(cfg.ServerConfig.ReadTimeoutSecs) * int64(time.Second)),
		WriteTimeout: time.Duration(int64(cfg.ServerConfig.WriteTimeoutSecs) * int64(time.Second)),
		Handler:      router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	graceful.Shutdown(srv, time.Duration(int64(cfg.ServerConfig.ShutdownTimeoutSecs)*int64(time.Second)))
}
