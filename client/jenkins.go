package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/theashishprasad/jenkins-build-report-tool/models"
)

// LoadBuild reads and parses build information from an HTTP endpoint.
func LoadBuild(url string) ([]models.Build, error) {
	var builds []models.Build

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return builds, err
	}

	jenkinsUser := os.Getenv("JENKINS_USER")
	jenkinsToken := os.Getenv("JENKINS_TOKEN")

	if jenkinsUser == "" || jenkinsToken == "" {
		return builds, fmt.Errorf("JENKINS_USER or JENKINS_TOKEN not set")
	}

	req.SetBasicAuth(jenkinsUser, jenkinsToken)

	resp, err := client.Do(req)
	if err != nil {
		return builds, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return builds, err
	}

	if resp.StatusCode > 299 {
		return builds, fmt.Errorf(
			"response failed with status code: %d and body: %s",
			resp.StatusCode,
			body,
		)
	}

	err = json.Unmarshal(body, &builds)
	if err != nil {
		return builds, err
	}

	return builds, nil
}
