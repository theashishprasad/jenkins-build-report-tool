# Jenkins Build Report Tool

A Go CLI application that fetches Jenkins build information from an HTTP endpoint and generates build statistics and status reports.

This project is part of my Go learning journey focused on DevOps, Platform Engineering, Cloud Infrastructure, and Automation.

## Project Goal

Build a Jenkins reporting tool while learning practical Go concepts used in real-world platform engineering and infrastructure tooling.

## Version 7 Features

* Fetch build information from an HTTP endpoint
* Support Jenkins authentication using environment variables
* Use HTTP Basic Authentication
* Accept endpoint URL as a command-line argument
* Parse JSON into Go structs and slices
* Support multiple build reports
* Generate build statistics and status aggregation
* Count builds by status
* Calculate total builds processed
* Handle HTTP request errors gracefully
* Handle HTTP response validation
* Handle JSON parsing errors gracefully
* Validate command-line input
* Organize code using reusable Go packages
* Separate application logic from business logic

## Learning Objectives

### Completed in Version 1

* Structs
* JSON
* Struct Tags
* json.Unmarshal
* File Reading
* Error Handling

### Completed in Version 2

* Packages
* Separation of Concerns
* Exported Functions
* Code Organization
* Reusable Components

### Completed in Version 3

* net/http
* HTTP Clients
* HTTP Requests
* HTTP Response Handling
* Response Body Processing
* API Communication
* Status Code Validation

### Completed in Version 4

* Conditional Logic
* Business Rules
* if / else if Statements
* Status Mapping
* Fallback Handling
* User-Friendly Reporting

### Completed in Version 5

* os.Args
* Command-Line Applications
* CLI Input Validation
* Runtime Configuration
* User Input Handling
* Parameterized API Requests

### Completed in Version 6

* Slices
* JSON Arrays
* Collection Processing
* Aggregation
* Statistics Generation
* map[string]int
* Counting and Grouping Data
* Iterating Over Collections

### Completed in Version 7

* os.Getenv()
* Environment Variables
* Basic Authentication
* HTTP Request Construction
* http.NewRequest()
* req.SetBasicAuth()
* Secure Configuration
* Runtime Credential Management

### Upcoming

* Real Jenkins Integration
* Advanced API Consumption
* Production-Grade Configuration

## Roadmap

### Version 1

Parse build information from a JSON file. ✅

### Version 2

Move logic into reusable packages. ✅

### Version 3

Fetch build data using HTTP APIs. ✅

### Version 4

Generate build status summaries. ✅

### Version 5

Accept input using command-line arguments. ✅

### Version 6

Support multiple build reports. ✅

### Version 7

Add Jenkins authentication. ✅

### Version 8

Integrate with a real Jenkins API.

## Project Structure

```text
jenkins-build-report-tool/
├── main.go
├── client/
│   └── jenkins.go
├── models/
│   └── build.go
├── sample/
│   └── build.json
├── README.md
├── .gitignore
└── go.mod
```

## How to Run

### Set Environment Variables

Linux/macOS:

```bash
export JENKINS_USER=test-user
export JENKINS_TOKEN=test-token
```

Windows PowerShell:

```powershell
$env:JENKINS_USER="test-user"
$env:JENKINS_TOKEN="test-token"
```

### Start Local HTTP Server

From the project root:

```bash
python3 -m http.server 8080
```

### Run Application

```bash
go run main.go http://localhost:8080/sample/build.json
```

## Sample API Response

```json
[
  {
    "name": "my-app-build",
    "number": 125,
    "result": "SUCCESS",
    "duration": 120000
  },
  {
    "name": "my-app-build",
    "number": 124,
    "result": "UNSTABLE",
    "duration": 210000
  },
  {
    "name": "my-app-build",
    "number": 123,
    "result": "FAILURE",
    "duration": 100000
  },
  {
    "name": "my-app-build",
    "number": 122,
    "result": "SUCCESS",
    "duration": 120000
  },
  {
    "name": "my-app-build",
    "number": 121,
    "result": "ABORTED",
    "duration": 150000
  }
]
```

## Sample Output

```text
Total Builds : 5

SUCCESS : 2
UNSTABLE : 1
FAILURE : 1
ABORTED : 1
```

> Note: Output order may vary because Go maps do not guarantee iteration order.

## Validation

```bash
go fmt ./...
go build ./...
```

### Manual Tests

* Valid credentials and valid endpoint
* Missing JENKINS_USER
* Missing JENKINS_TOKEN
* Invalid credentials
* Missing CLI argument
* Invalid URL
* HTTP server unavailable
* Invalid JSON response
* Non-200 HTTP response
* Multiple build records
* Empty JSON array
* Mixed build statuses

## Environment Variables

| Variable      | Description                   |
| ------------- | ----------------------------- |
| JENKINS_USER  | Jenkins username              |
| JENKINS_TOKEN | Jenkins API token or password |

## Version

Current Version:

```text
v0.7.0
```
