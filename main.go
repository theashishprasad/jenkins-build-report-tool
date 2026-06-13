package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/theashishprasad/jenkins-build-report-tool/models"
)

func main() {
	jsonData, err := os.ReadFile("sample/build.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var build models.Build

	err = json.Unmarshal(jsonData, &build)
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
}