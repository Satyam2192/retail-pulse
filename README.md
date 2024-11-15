# Retail Pulse Image Processing Service

A scalable service for processing retail store images, built with Go.

## Features

- Asynchronous image processing
- Concurrent job handling
- RESTful API endpoints
- Store management system
- Job status tracking
- Error handling and logging

## Prerequisites

- Go 1.21 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/satyam2192/retail-pulse.git
cd retail-pulse
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o retail-pulse ./cmd/server
```
4. Start the server:
```bash
./retail-pulse
```

## Withh Docker

1. Build and Start the Application
```bash
docker-compose up --build
```
## Testing

Run the tests:
```bash
go test ./...
```

The service will start on port 7000 by default.

## Requirements:
=> Job Submission and Processing:

✅ Implements POST /api/submit endpoint
✅ Handles multiple jobs concurrently using worker pool
✅ Validates job payload (count, visits, store IDs)
✅ Implements image download and perimeter calculation
✅ Includes random sleep time (0.1-0.4s) to simulate GPU processing

=> Store Master Integration:

✅ Loads store data from CSV
✅ Validates store IDs against master data
✅ Provides store lookup functionality


=> Job Status Tracking:

✅ Implements GET /api/status endpoint
✅ Tracks job status (ongoing/completed/failed)
✅ Handles errors appropriately
✅ Returns proper error responses

=> Technical Requirements:

✅ Written in Go with Go Modules
✅ Uses Gin framework for routing
✅ Implements proper error handling
✅ Includes logging
✅ Uses goroutines for concurrent processing

## API Endpoints

### Submit Job
- **URL**: `/api/submit`
- **Method**: `POST`
- **Request Body**:
```json
{
   "count": 2,
   "visits": [
      {
         "store_id": "S00339218",
         "image_url": [
            "https://example.com/image1.jpg",
            "https://example.com/image2.jpg"
         ],
         "visit_time": "2024-03-15T10:00:00Z"
      }
   ]
}
```

### Check Job Status
- **URL**: `/api/status`
- **Method**: `GET`
- **Query Parameters**: `jobid`


## Configuration

Configuration can be modified in `internal/config/config.go`:
- Server address
- Maximum workers
- Queue size
- Other settings


## License

This project is licensed under the MIT License - see the LICENSE file for details.
