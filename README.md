# Retail Pulse Image Processing Service

- Postman: https://documenter.getpostman.com/view/31555061/2sAYBPkDkn
- Live: https://retail-pulse.onrender.com/api/status?jobid=1731677266939497569

## Description
The Retail Pulse Image Processing Service is a scalable Go application that allows users to submit jobs to process images collected from retail stores. The service downloads the images, calculates the perimeter of each image, and stores the results. It also handles errors and provides job status tracking

## Features

- Asynchronous image processing
- Concurrent job handling
- RESTful API endpoints
- Store management system
- Job status tracking
- Error handling and logging

## Assumptions

1. The service assumes that the input data (job payload) is valid and follows the specified format.
2. The service assumes that the store master data in the CSV file is correct and up-to-date.
3. The service assumes that the image URLs are accessible and the images can be successfully downloaded.
4. The service assumes that the image processing (perimeter calculation) can be completed within a reasonable time frame.

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

The service will start on port 7000 by default.

## Requirements Fulfilled:
=> Job Submission and Processing:

- ✅ Implements POST /api/submit endpoint
- ✅ Handles multiple jobs concurrently using worker pool
- ✅ Validates job payload (count, visits, store IDs)
- ✅ Implements image download and perimeter calculation
- ✅ Includes random sleep time (0.1-0.4s) to simulate GPU processing

=> Store Master Integration:

- ✅ Loads store data from CSV
- ✅ Validates store IDs against master data
- ✅ Provides store lookup functionality


=> Job Status Tracking:

- ✅ Implements GET /api/status endpoint
- ✅ Tracks job status (ongoing/completed/failed)
- ✅ Handles errors appropriately
- ✅ Returns proper error responses

=> Technical Requirements:

- ✅ Written in Go with Go Modules
- ✅ Uses Gin framework for routing
- ✅ Implements proper error handling
- ✅ Includes logging
- ✅ Uses goroutines for concurrent processing

## Work Environment

- Computer/Operating System: ubuntu 24.04
- Text Editor: Visual Studio Code 1.74.2

=> Libraries used:

- Gin-Gonic/Gin: web framework
- Encoding/csv: for reading CSV files
- Image/image, Image/jpeg, Image/png: for image processing
- Math/rand: for generating random sleep durations
- Sync: for concurrency primitives
- Log: for logging

## API Endpoints

### Submit Job
- **URL**: `https://retail-pulse.onrender.com/api/submit`
- **Method**: `POST`
- **Request Body**:
```json
{
   "count": 2,
   "visits": [
      {
         "store_id": "RP00001",
         "image_urls": [
            "https://www.gstatic.com/webp/gallery/2.jpg",
            "https://www.gstatic.com/webp/gallery/3.jpg"
         ],
         "visit_time": "2024-11-15T10:00:00Z"
      },
      {
         "store_id": "RP00002",
         "image_urls": [
            "https://www.gstatic.com/webp/gallery/4.jpg"
         ],
         "visit_time": "2024-11-15T11:00:00Z"
      }
   ]
}
```

### Check Job Status
- **URL**: `https://retail-pulse.onrender.com/api/status?jobid={{jobid}}`
- **Example**: `https://retail-pulse.onrender.com/api/status?jobid=1731677266939497569`
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
