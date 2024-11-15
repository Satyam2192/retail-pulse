package main

import (
    "log"
    "retail-pulse/internal/api/router"
    "retail-pulse/internal/config"
    "retail-pulse/internal/store"
    "retail-pulse/pkg/logger"
	"os"
)

func main() {
    // Initialize logger
    logger := logger.New()

    // Load configuration
    cfg, err := config.Load()
    if err != nil {
        logger.Fatal("Failed to load configuration:", err)
    }

    // Print current working directory
    pwd, err := os.Getwd()
    if err != nil {
        logger.Fatal("Failed to get working directory:", err)
    }
    logger.Infof("Current working directory: %s", pwd)
    logger.Infof("Loading stores from: %s", cfg.StoreMasterPath)

    // Load stores from CSV
    err = store.LoadStoresFromCSV(cfg.StoreMasterPath)
    if err != nil {
        logger.Fatal("Failed to load store master data:", err)
    }

    // Initialize and start the server
    r := router.Setup(logger)
    logger.Infof("Server starting on %s", cfg.ServerAddress)
    log.Fatal(r.Run(cfg.ServerAddress))
}