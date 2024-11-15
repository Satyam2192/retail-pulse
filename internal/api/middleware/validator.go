package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "retail-pulse/internal/models"
)

func ValidateJobMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.Request.Method == "POST" && c.Request.URL.Path == "/api/submit" {
            var job models.Job
            if err := c.ShouldBindJSON(&job); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload: " + err.Error()})
                c.Abort()
                return
            }

            if job.Count <= 0 {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Count must be greater than 0"})
                c.Abort()
                return
            }

            if len(job.Visits) != job.Count {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Count does not match number of visits"})
                c.Abort()
                return
            }

            for _, visit := range job.Visits {
                if visit.StoreID == "" {
                    c.JSON(http.StatusBadRequest, gin.H{"error": "Store ID cannot be empty"})
                    c.Abort()
                    return
                }
                if len(visit.ImageURLs) == 0 {
                    c.JSON(http.StatusBadRequest, gin.H{"error": "Image URLs cannot be empty"})
                    c.Abort()
                    return
                }
            }

            c.Set("validatedJob", job)
        }
        c.Next()
    }
}