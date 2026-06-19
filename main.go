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

	builds, err := client.LoadBuild(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalBuilds := len(builds)
	fmt.Printf("Total Builds : %d\n\n", totalBuilds)

	status := make(map[string]int)

	for _, build := range builds {
		status[build.Result]++
	}

	for key, val := range status {
		fmt.Printf("%s : %d\n", key, val)
	}

}
