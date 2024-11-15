package processor

import (
    "retail-pulse/internal/models"
    "retail-pulse/internal/store"
    "retail-pulse/pkg/logger"
)

type ImageProcessor struct {
    resultChan chan models.Result
    errorChan  chan models.Error
    logger     *logger.Logger
}

func NewImageProcessor(logger *logger.Logger) *ImageProcessor {
    return &ImageProcessor{
        resultChan: make(chan models.Result),
        errorChan:  make(chan models.Error),
        logger:     logger,
    }
}

func (p *ImageProcessor) ProcessJob(job *models.Job) {
    workerPool := NewWorkerPool(5, 100, p.logger) // Use config values
    
    totalImages := 0
    for _, visit := range job.Visits {
        totalImages += len(visit.ImageURLs)
    }

    results := make(chan models.Result, totalImages)
    errors := make(chan models.Error, totalImages)
    
    // Submit work items
    for _, visit := range job.Visits {
        if _, exists := store.GetStore(visit.StoreID); !exists {
            job.Errors = append(job.Errors, models.Error{
                StoreID: visit.StoreID,
                Message: "Store not found",
            })
            job.Status = "failed"
            return
        }

        for _, imageURL := range visit.ImageURLs {
            workerPool.jobQueue <- &WorkItem{
                storeID:    visit.StoreID,
                imageURL:   imageURL,
                resultChan: results,
                errorChan:  errors,
            }
        }
    }

    // Collect results
    processed := 0
    for processed < totalImages {
        select {
        case result := <-results:
            job.Results = append(job.Results, result)
        case err := <-errors:
            job.Errors = append(job.Errors, err)
            job.Status = "failed"
            return
        }
        processed++
    }

    if len(job.Errors) == 0 {
        job.Status = "completed"
    }
}