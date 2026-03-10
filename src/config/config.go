package config

import (
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	Port        int
	Environment string
	ContentDir  string
	StaticDir   string
}

// Load reads configuration from environment variables with sensible defaults
func Load() *Config {
	port := 3000
	if portStr := os.Getenv("PORT"); portStr != "" {
		if p, err := strconv.Atoi(portStr); err == nil {
			port = p
		}
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	contentDir := os.Getenv("CONTENT_DIR")
	if contentDir == "" {
		contentDir = "./content/blog"
	}

	staticDir := os.Getenv("STATIC_DIR")
	if staticDir == "" {
		staticDir = "./src/static"
	}

	return &Config{
		Port:        port,
		Environment: env,
		ContentDir:  contentDir,
		StaticDir:   staticDir,
	}
}
