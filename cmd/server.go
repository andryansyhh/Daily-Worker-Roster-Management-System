package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"worker-management/internal/router"

	"github.com/gin-gonic/gin"
)

func Run() {
	deps := InitDependencies()

	r := gin.Default()
	router.SetupRoutes(r, deps.ShiftHandler, deps.ShiftReqHandler, deps.AssignmentReqHandler, *deps.AuthHandler, deps.WorkerHandler)

	server := &http.Server{
		Addr:    ":" + deps.Config.Port,
		Handler: r,
	}

	go func() {
		log.Printf("HTTP server listening on port %s", deps.Config.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down HTTP server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited gracefully.")
}
