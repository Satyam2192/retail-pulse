package processor

import (
    "bytes"
    "fmt"
    "image"
    _ "image/jpeg"
    _ "image/png"
    "io"
    "math/rand"
    "net/http"
    "time"
    "retail-pulse/internal/models"
    "retail-pulse/pkg/logger"
)

type WorkerPool struct {
    maxWorkers   int
    maxQueueSize int
    jobQueue     chan *WorkItem
    logger       *logger.Logger
}

type WorkItem struct {
    storeID    string
    imageURL   string
    resultChan chan<- models.Result
    errorChan  chan<- models.Error
}

func NewWorkerPool(maxWorkers, maxQueueSize int, logger *logger.Logger) *WorkerPool {
    pool := &WorkerPool{
        maxWorkers:   maxWorkers,
        maxQueueSize: maxQueueSize,
        jobQueue:     make(chan *WorkItem, maxQueueSize),
        logger:       logger,
    }
    
    pool.Start()
    return pool
}

func (p *WorkerPool) Start() {
    for i := 0; i < p.maxWorkers; i++ {
        go p.worker()
    }
}

func (p *WorkerPool) worker() {
    for item := range p.jobQueue {
        perimeter, err := downloadAndCalculatePerimeter(item.imageURL)
        if err != nil {
            item.errorChan <- models.Error{
                StoreID: item.storeID,
                Message: err.Error(),
            }
            continue
        }

        // Simulate processing time (0.1 to 0.4 seconds)
        time.Sleep(time.Duration(100+rand.Intn(300)) * time.Millisecond)

        item.resultChan <- models.Result{
            StoreID:   item.storeID,
            ImageURL:  item.imageURL,
            Perimeter: perimeter,
        }
    }
}

// downloadAndCalculatePerimeter downloads an image from URL and calculates its perimeter
func downloadAndCalculatePerimeter(imageURL string) (float64, error) {
    resp, err := http.Get(imageURL)
    if err != nil {
        return 0, fmt.Errorf("failed to download image: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return 0, fmt.Errorf("failed to download image, status: %d", resp.StatusCode)
    }

    img, _, err := image.Decode(resp.Body)
    if err != nil {
        // If decode fails, try to read the full body and decode again
        // This handles some edge cases with image formats
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            return 0, fmt.Errorf("failed to read image body: %v", err)
        }
        
        img, _, err = image.Decode(bytes.NewReader(body))
        if err != nil {
            return 0, fmt.Errorf("failed to decode image: %v", err)
        }
    }

    // Calculate perimeter
    bounds := img.Bounds()
    height := bounds.Max.Y - bounds.Min.Y
    width := bounds.Max.X - bounds.Min.X

    // Perimeter = 2 * (height + width)
    perimeter := float64(2 * (height + width))

    return perimeter, nil
}