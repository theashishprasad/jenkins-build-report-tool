package main

import (
	"fmt"
	"os"

	"github.com/theashishprasad/jenkins-build-report-tool/client"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run main.go <endpoint-url>")
		return
	}

	url := args[1]

	build, err := client.LoadBuild(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Build Report")
	fmt.Println()

	fmt.Printf("Job Name : %s\n", build.Name)
	fmt.Printf("Build No : %d\n", build.Number)
	fmt.Printf("Status   : %s\n", build.Result)
	fmt.Printf("Duration : %d sec\n", build.Duration/1000)

	summary := "Unknown build status."

	fmt.Println("\nSummary\n")

	if build.Result == "SUCCESS" {
		summary = "Build completed successfully."
	} else if build.Result == "FAILURE" {
		summary = "Build failed. Investigation required."
	} else if build.Result == "ABORTED" {
		summary = "Build was aborted."
	} else if build.Result == "UNSTABLE" {
		summary = "Build completed with warnings."
	}

	fmt.Println(summary)

}
