# Jenkins Build Report Tool

A Go CLI application that reads Jenkins build information from a JSON file and generates a simple build report.

This project is part of my Go learning journey focused on DevOps, Platform Engineering, Cloud Infrastructure, and Automation.

## Project Goal

Build a Jenkins reporting tool while learning practical Go concepts used in real-world platform engineering and infrastructure tooling.

## Version 2 Features

* Read build information from a JSON file
* Parse JSON into Go structs
* Generate a formatted build report
* Handle file reading errors gracefully
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

### Upcoming

* HTTP clients and REST APIs
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

Fetch build data using HTTP APIs.

### Version 4

Generate build status summaries.

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

Run the application:

```bash
go run main.go
```

## Sample Input

```json
{
  "name": "my-app-build",
  "number": 125,
  "result": "SUCCESS",
  "duration": 120000
}
```

## Sample Output

```text
Build Report

Job Name : my-app-build
Build No : 125
Status   : SUCCESS
Duration : 120 sec
```

## Validation

```bash
go fmt ./...
go build ./...
go run main.go
```

## Version

Current Version:

```text
v0.2.0
```
