package main

import (
	"fmt"
	"os"
	"strings"

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

	fmt.Println("Build Report\n")

	fmt.Printf("Job Name : %s\n", strings.Split(build.FullDisplayName, " #")[0])
	fmt.Printf("Build No : %d\n", build.Number)
	fmt.Printf("Status   : %s\n", build.Result)
	fmt.Printf("Duration : %.2f sec\n", float64(build.Duration)/1000)

}
