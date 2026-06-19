package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/theashishprasad/jenkins-build-report-tool/models"
)

// LoadBuild reads and parses build information from an HTTP endpoint.
func LoadBuild(url string) ([]models.Build, error) {
	var builds []models.Build

	res, err := http.Get(url)
	if err != nil {
		return builds, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return builds, err
	}

	if res.StatusCode > 299 {
		return builds, fmt.Errorf(
			"response failed with status code: %d and body: %s",
			res.StatusCode,
			body,
		)
	}

	err = json.Unmarshal(body, &builds)
	if err != nil {
		return builds, err
	}

	return builds, nil
}
