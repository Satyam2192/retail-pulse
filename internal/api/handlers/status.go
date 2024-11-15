package handlers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

func (h *Handler) HandleStatus(c *gin.Context) {
    jobID, err := strconv.ParseInt(c.Query("jobid"), 10, 64)
    if err != nil {
        h.logger.Errorf("Invalid job ID: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
        return
    }

    job, exists := h.jobService.GetJob(jobID)
    if !exists {
        h.logger.Infof("Job not found: %d", jobID)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Job not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": job.Status,
        "job_id": job.ID,
        "error":  job.Errors,
    })
}
