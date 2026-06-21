# Jenkins Build Report Tool

A Go CLI application that integrates with the Jenkins API to retrieve build information and generate build reports.

This project is part of my Go learning journey focused on DevOps, Platform Engineering, Cloud Infrastructure, and Automation.

## Project Goal

Build a Jenkins reporting tool while learning practical Go concepts used in real-world platform engineering and infrastructure tooling.

## Final Features

* Integrate with the Jenkins REST API
* Support Jenkins authentication using environment variables
* Use HTTP Basic Authentication
* Accept Jenkins API endpoint URL as a command-line argument
* Parse Jenkins API responses into Go structs
* Generate formatted build reports
* Handle HTTP request errors gracefully
* Handle HTTP response validation
* Handle JSON parsing errors gracefully
* Validate command-line input
* Organize code using reusable Go packages
* Separate application logic from business logic

## Learning Objectives

### Version 1

* Structs
* JSON
* Struct Tags
* json.Unmarshal
* File Reading
* Error Handling

### Version 2

* Packages
* Separation of Concerns
* Exported Functions
* Code Organization
* Reusable Components

### Version 3

* net/http
* HTTP Clients
* HTTP Requests
* HTTP Response Handling
* Response Body Processing
* API Communication
* Status Code Validation

### Version 4

* Conditional Logic
* Business Rules
* if / else if Statements
* Status Mapping
* Fallback Handling
* User-Friendly Reporting

### Version 5

* os.Args
* Command-Line Applications
* CLI Input Validation
* Runtime Configuration
* User Input Handling
* Parameterized API Requests

### Version 6

* Slices
* JSON Arrays
* Collection Processing
* Aggregation
* Statistics Generation
* map[string]int
* Counting and Grouping Data
* Iterating Over Collections

### Version 7

* os.Getenv()
* Environment Variables
* Basic Authentication
* HTTP Request Construction
* http.NewRequest()
* req.SetBasicAuth()
* Secure Configuration
* Runtime Credential Management

### Version 8

* Jenkins API Integration
* API Contract Mapping
* Production-Style Tooling
* External System Integration
* Real API Consumption
* Build Metadata Processing

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

Integrate with a real Jenkins API. ✅

## Project Structure

```text
jenkins-build-report-tool/
├── main.go
├── client/
│   └── jenkins.go
├── models/
│   └── build.go
├── README.md
├── .gitignore
└── go.mod
```

## How to Run

### Set Environment Variables

Linux/macOS:

```bash
export JENKINS_USER=my-user
export JENKINS_TOKEN=my-token
```

Windows PowerShell:

```powershell
$env:JENKINS_USER="my-user"
$env:JENKINS_TOKEN="my-token"
```

### Run Application

```bash
go run main.go "<jenkins-api-url>"
```

Example:

```bash
go run main.go "https://jenkins.example.com/job/service-a/456/api/json"
```

## Example Jenkins API Response

```json
{
  "fullDisplayName": "service-a #456",
  "number": 456,
  "result": "SUCCESS",
  "duration": 180000
}
```

## Sample Output

```text
Build Report

Job Name : service-a
Build No : 456
Status   : SUCCESS
Duration : 180.00 sec
```

## Validation

```bash
go fmt ./...
go build ./...
```

## Manual Tests

* Valid Jenkins credentials
* Missing JENKINS_USER
* Missing JENKINS_TOKEN
* Invalid credentials (401)
* Missing CLI argument
* Invalid URL
* Jenkins unavailable
* Invalid JSON response
* Non-200 HTTP response
* Successful Jenkins API response

## Environment Variables

| Variable      | Description                   |
| ------------- | ----------------------------- |
| JENKINS_USER  | Jenkins username              |
| JENKINS_TOKEN | Jenkins API token or password |

## Technologies Used

* Go
* net/http
* JSON
* Basic Authentication
* Environment Variables
* CLI Arguments
* REST APIs

## Learning Outcomes

Through this project I practiced:

* Building CLI applications in Go
* Designing reusable packages
* Working with JSON APIs
* Consuming REST endpoints
* Implementing authentication
* Managing runtime configuration
* Handling errors idiomatically
* Integrating with external systems
* Building production-style infrastructure tooling

## Version

```text
v0.8.0
```
