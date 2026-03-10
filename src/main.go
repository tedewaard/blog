package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tedewaard/folio/src/config"
	"github.com/tedewaard/folio/src/handlers"
	"github.com/tedewaard/folio/src/middleware"
	"github.com/tedewaard/folio/src/services"
)

func main() {
	// Load configuration
	cfg := config.Load()
	log.Printf("Starting server on port %d in %s mode", cfg.Port, cfg.Environment)

	// Initialize services
	blogService := services.NewBlogService(cfg.ContentDir)

	// Initialize handlers
	blogHandler := handlers.NewBlogHandler(blogService)

	// Set up router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Recovery)
	r.Use(middleware.Logger)

	// Routes
	r.Get("/", handlers.HandleHome)
	r.Get("/blog", blogHandler.HandleList)
	r.Get("/blog/{slug}", blogHandler.HandlePost)

	// Static files
	fileServer := http.FileServer(http.Dir(cfg.StaticDir))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Create server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server listening on http://localhost:%d", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}
