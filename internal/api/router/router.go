package router

import (
    "github.com/gin-gonic/gin"
    "retail-pulse/internal/api/handlers"
    "retail-pulse/pkg/logger"
    "retail-pulse/internal/service"
)

func Setup(logger *logger.Logger) *gin.Engine {
    r := gin.New()
    
    // Add middlewares
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    
    // Configure to parse JSON payload
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Content-Type", "application/json")
        c.Next()
    })

    jobService := service.NewJobService()
    handler := handlers.NewHandler(jobService, logger)

    // Setup routes
    api := r.Group("/api")
    {
        api.POST("/submit", handler.HandleSubmit)
        api.GET("/status", handler.HandleStatus)
    }

    return r
}