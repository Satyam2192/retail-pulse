// internal/api/handlers/handlers.go
package handlers

import (
    "errors"
    "time"
    "github.com/gin-gonic/gin"
    "retail-pulse/internal/models"
    "retail-pulse/internal/service"
    "retail-pulse/pkg/logger"
)


type Handler struct {
    jobService *service.JobService
    logger     *logger.Logger
}

func NewHandler(jobService *service.JobService, logger *logger.Logger) *Handler {
    return &Handler{
        jobService: jobService,
        logger:     logger,
    }
}

// Response represents a standard API response
type Response struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string     `json:"error,omitempty"`
}

// newResponse creates a new Response
func newResponse(success bool, data interface{}, err string) Response {
    return Response{
        Success: success,
        Data:    data,
        Error:   err,
    }
}

// validateJob validates the job request
func (h *Handler) validateJob(job *models.Job) error {
    if job.Count <= 0 {
        return errors.New("count must be greater than 0")
    }

    if len(job.Visits) == 0 {
        return errors.New("visits cannot be empty")
    }

    if job.Count != len(job.Visits) {
        return errors.New("count does not match number of visits")
    }

    for _, visit := range job.Visits {
        if visit.StoreID == "" {
            return errors.New("store_id cannot be empty")
        }
        if len(visit.ImageURLs) == 0 {
            return errors.New("image_urls cannot be empty")
        }
        for _, url := range visit.ImageURLs {
            if url == "" {
                return errors.New("image url cannot be empty")
            }
            // Add more URL validation if needed
        }
    }

    return nil
}

// Middleware for logging
func (h *Handler) LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Start timer
        start := time.Now()

        // Process request
        c.Next()

        // Log request details
        duration := time.Since(start)
        h.logger.Infof(
            "Method: %s | URL: %s | Status: %d | Duration: %v",
            c.Request.Method,
            c.Request.URL.Path,
            c.Writer.Status(),
            duration,
        )
    }
}

// Middleware for error handling
func (h *Handler) ErrorMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

        // Check if there are any errors
        if len(c.Errors) > 0 {
            c.JSON(c.Writer.Status(), newResponse(false, nil, c.Errors.Last().Error()))
            return
        }
    }
}