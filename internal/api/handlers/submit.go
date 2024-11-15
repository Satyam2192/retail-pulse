package handlers

import (
    "net/http"
    "time"
    "bytes"
    "io"
    "github.com/gin-gonic/gin"
    "retail-pulse/internal/models"
    "retail-pulse/internal/processor"
    "retail-pulse/internal/store"  // Add this import
)

func (h *Handler) HandleSubmit(c *gin.Context) {
    // Read the request body
    bodyBytes, err := io.ReadAll(c.Request.Body)
    if err != nil {
        h.logger.Errorf("Failed to read request body: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
        return
    }

    c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
    
    h.logger.Infof("Raw request body: %s", string(bodyBytes))

    var job models.Job
    if err := c.BindJSON(&job); err != nil {
        h.logger.Errorf("Failed to bind JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format: " + err.Error()})
        return
    }

    // Basic validation
    if job.Count <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Count must be greater than 0"})
        return
    }

    if len(job.Visits) != job.Count {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Count does not match number of visits"})
        return
    }

    // Validate store IDs first
    for _, visit := range job.Visits {
        if visit.StoreID == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Store ID cannot be empty"})
            return
        }
        // Check if store exists
        if _, exists := store.GetStore(visit.StoreID); !exists {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Store not found"})
            return
        }
    }

    // Then validate other fields
    for _, visit := range job.Visits {
        if len(visit.ImageURLs) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Image URLs cannot be empty"})
            return
        }
    }

    // Generate job ID and set initial status
    job.ID = time.Now().UnixNano()
    job.Status = "ongoing"

    // Store job
    h.jobService.CreateJob(&job)

    // Process job asynchronously
    go processor.NewImageProcessor(h.logger).ProcessJob(&job)

    c.JSON(http.StatusCreated, gin.H{"job_id": job.ID})
}