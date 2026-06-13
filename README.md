# Jenkins Build Report Tool

A Go CLI application that retrieves Jenkins build information and generates build reports.

This project is part of my Go learning journey focused on DevOps, Platform Engineering, Cloud Infrastructure, and Automation.

## Project Goal

Build a Jenkins reporting tool while learning practical Go concepts used in real-world platform engineering and infrastructure tooling.

## Learning Objectives

* Structs and JSON parsing
* HTTP clients and REST APIs
* Packages and project organization
* Error handling
* Command-line applications
* Environment variables
* Authentication
* Working with external systems

## Roadmap

### Version 1

Parse build information from a JSON file.

### Version 2

Move logic into reusable packages.

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