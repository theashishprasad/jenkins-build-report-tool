# Jenkins Build Report Tool

A Go CLI application that fetches Jenkins build information from an HTTP endpoint and generates a simple build report.

This project is part of my Go learning journey focused on DevOps, Platform Engineering, Cloud Infrastructure, and Automation.

## Project Goal

Build a Jenkins reporting tool while learning practical Go concepts used in real-world platform engineering and infrastructure tooling.

## Version 4 Features

* Fetch build information from an HTTP endpoint
* Parse JSON into Go structs
* Generate a formatted build report
* Generate human-readable build status summaries
* Handle HTTP request errors gracefully
* Handle HTTP response validation
* Handle JSON parsing errors gracefully
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

### Upcoming

* Command-line applications
* Environment variables
* Authentication
* Working with external systems

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

Accept input using command-line arguments.

### Version 6

Support multiple build reports.

### Version 7

Add Jenkins authentication.

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

### Start Local HTTP Server

From the project root:

```bash
python3 -m http.server 8080
```

### Run Application

In a separate terminal:

```bash
go run main.go
```

## Sample API Response

```json
{
  "name": "my-app-build",
  "number": 125,
  "result": "SUCCESS",
  "duration": 120000
}
```

Endpoint:

```text
http://localhost:8080/sample/build.json
```

## Sample Output

```text
Build Report

Job Name : my-app-build
Build No : 125
Status   : SUCCESS
Duration : 120 sec

Summary

Build completed successfully.
```

## Supported Build Statuses

| Status          | Summary                               |
| --------------- | ------------------------------------- |
| SUCCESS         | Build completed successfully.         |
| FAILURE         | Build failed. Investigation required. |
| ABORTED         | Build was aborted.                    |
| UNSTABLE        | Build completed with warnings.        |
| Any Other Value | Unknown build status.                 |

## Validation

```bash
go fmt ./...
go build ./...
go run main.go
```

### Manual Tests

* Valid HTTP response
* HTTP server unavailable
* Invalid JSON response
* Non-200 HTTP response
* SUCCESS status
* FAILURE status
* ABORTED status
* UNSTABLE status
* Unknown status

## Version

Current Version:

```text
v0.4.0
```
