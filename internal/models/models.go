package models

type Job struct {
    ID      int64     `json:"job_id"`
    Status  string    `json:"status"`
    Visits  []Visit   `json:"visits,omitempty"`
    Count   int       `json:"count,omitempty"`
    Results []Result  `json:"results,omitempty"`
    Errors  []Error   `json:"error,omitempty"`
}

type Visit struct {
    StoreID   string   `json:"store_id"`
    ImageURLs []string `json:"image_urls"` 
    VisitTime string   `json:"visit_time"`
}

type Result struct {
    StoreID   string  `json:"store_id"`
    ImageURL  string  `json:"image_urls"`
    Perimeter float64 `json:"perimeter"`
}

type Error struct {
    StoreID string `json:"store_id"`
    Message string `json:"error"`
}

type Store struct {
    AreaCode  string
    StoreName string
    StoreID   string
}