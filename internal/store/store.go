package store

import (
    "encoding/csv"
    "fmt"
    "os"
    "path/filepath"
    "sync"
    "retail-pulse/internal/models"
)

var (
    stores   = make(map[string]models.Store)
    storeMux sync.RWMutex
)

func LoadStoresFromCSV(path string) error {
    // Get absolute path
    absPath, err := filepath.Abs(path)
    if err != nil {
        return fmt.Errorf("failed to get absolute path: %v", err)
    }

    // Check if file exists
    if _, err := os.Stat(absPath); os.IsNotExist(err) {
        return fmt.Errorf("file does not exist at path: %s", absPath)
    }

    file, err := os.Open(absPath)
    if err != nil {
        return fmt.Errorf("failed to open file: %v", err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    // Skip header
    _, err = reader.Read()
    if err != nil {
        return fmt.Errorf("failed to read header: %v", err)
    }

    storeMux.Lock()
    defer storeMux.Unlock()

    stores = make(map[string]models.Store)

    for {
        record, err := reader.Read()
        if err != nil {
            if err.Error() == "EOF" {
                break // Normal EOF, we're done reading
            }
            return fmt.Errorf("error reading CSV: %v", err)
        }

        if len(record) >= 3 {
            stores[record[2]] = models.Store{
                AreaCode:  record[0],
                StoreName: record[1],
                StoreID:   record[2],
            }
        }
    }

    // Add debug print
    fmt.Printf("Loaded %d stores from CSV\n", len(stores))
    return nil
}


func GetStore(id string) (models.Store, bool) {
    storeMux.RLock()
    defer storeMux.RUnlock()
    
    store, exists := stores[id]
    return store, exists
}


