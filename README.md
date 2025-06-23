# Data Ingestion Project 

This service is a Go-based application built with the Gin web framework. It provides a set of RESTful API endpoints to ingest and query time-series data, specifically CPU load and concurrency metrics, stored in a PostgreSQL database.

**The primary functionalities include:**

- Ingesting data sets for the past five minutes.
- Querying data based on various time intervals (days, hours, minutes, seconds).
- Retrieving data for a specific calendar date.

---


##  Project Structure

- `cmd/server/main.go`: The main entry point for starting the backend API server.
- `cmd/script/main.go`: Entry point for a standalone script that ingests data from the last five minutes into the database.
- `cmd/server/service`: Implements the core business logic and handlers for the API endpoints.
- `cmd/server/models`: Contains the Go struct definitions for API request and response bodies.
- `cmd/server/db`: Contains the Go struct definitions that map directly to the database table schema.
- `cmd/server/utils`: Provides shared utility functions, such as query helpers, used across the service layer.

---
## Prerequisite
- Golang - Core language for the application backend.
- Gin - High-performance HTTP web framework for Go.
- TimescaleDb - PostgreSQL extension for time-series data, providing performance optimizations and lifecycle management.

### Env Variables
- DATABASE_URL = CONNECTION_STRING 
    -   Example : `export DATABASE_URL=postgres://postgres:password@localhost:5432/postgres`

### DB Table Creation
- Excute the sql script in `/resources/01_CreateTable.sql` which will create an table with required feilds.
---

## How to Run

### 1. Script - Ingest data to DB 
- Navigate to `cmd/script`
- Execute `go run main.go`

### 2. Server
- Navigate to `cmd/server`
- Execute `go run main.go`
- Rest server will be available on port number: 8080 
- Use Postman or REST Client to test

---
##  API Endpoints

#### **Get Data for the Last N Days**

- **Endpoint:** GET `/days/:day`
- **Description:** Retrieves all data points from the last N days, where N is the specified day parameter.
- **URL Params:**
  - day=[integer] **(Required)**: The number of past days to fetch data for.

- **Success Response (200 OK):**
```
{
    "data": [
    
        {
        "timestamp": 1719139800,
        "cpu_load": 50.5,
        "concurrency": 12345
        },
        
        {
        "timestamp": 1719136200,
        "cpu_load": 75.0,
        "concurrency": 54321
        }
    ]
}
```
* **Error Response (400 Bad Request):**
  ```
  {
    "error": "Invalid day parameter, must be an integer greater than zero"
  }
  ```



---


#### Get Data for the Last N Hours

- **Endpoint:** GET `/hours/:hour`
- **Description:** Retrieves all data points from the last N hours.
- **URL Params:**
  * hour=[integer] (Required)
- **Example Request:** GET /hours/12


---


#### Get Data for the Last N Seconds

* **Endpoint:** GET `/seconds/:second`
* **Description:** Retrieves all data points from the last N seconds.
* **URL Params:**
  * second=[integer] (Required)
* **Example Request:** GET /seconds/45


---


#### Get Data by Specific Date

* **Endpoint:** GET `/date/:date`
* **Description:** Retrieves all data points for a specific calendar date.
* **URL Params:**
  * date=[string] (Required): The date in DD-MM-YYYY format.
* **Example Request:** GET /date/23-06-2025


---


#### Get Specific or Aggregated Data Set

* **Endpoint:** POST `/getData`
* **Description:** This endpoint has dual functionality. It can either fetch raw data points within a specific time range or perform an aggregate calculation (avg, min, max) on a specified data field (cpu_load or concurrency) within that range.
* **Request Body:**

  * **starttime (string, Required):** The start of the time range in RFC3339 format (e.g., "2025-06-23T10:00:00Z").
  * **endtime (string, Required):** The end of the time range in RFC3339 format.
  * **opcode (string, Optional):** The aggregate operation to perform. Supported values: **"avg"**, **"max"**. If omitted, raw data is returned.
  * **params (string, Optional):** The data field to perform the aggregation on. Required if opcode is present. Supported values: **"cpu_load"**, **"concurrency"**


* **Example Request (Raw Data):**
  ```
  {
      "starttime": "2025-06-23T00:00:00Z",
      "endtime": "2025-06-23T23:59:59Z"
  }
  ```



* **Example Request (Aggregated Data):**
 ```
    {
          "starttime": "2025-06-23T00:00:00Z",
          "endtime": "2025-06-23T23:59:59Z",
          "opcode": "avg",
          "params": "cpu_load"
    }
```

### Future Scope
- Testing - Test case implemention using Mock.
- Code Cleanup and Scope for Refactoring.
- Isolate the the DB operation to seperate module, to support other databases.