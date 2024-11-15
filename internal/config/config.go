package config

type Config struct {
    ServerAddress  string
    MaxWorkers     int
    MaxQueueSize   int
    StoreMasterPath string
}

func Load() (*Config, error) {
    return &Config{
        ServerAddress:   ":7000",
        MaxWorkers:      5,
        MaxQueueSize:    100,
        StoreMasterPath: "./StoreMasterAssignment.csv", 
    }, nil
}