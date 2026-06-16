package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/theashishprasad/jenkins-build-report-tool/models"
)

// LoadBuild reads and parses build information from an HTTP endpoint.
func LoadBuild() (models.Build, error) {
	var build models.Build

	res, err := http.Get("http://localhost:8080/sample/build.json")
	if err != nil {
		return build, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return build, err
	}

	if res.StatusCode > 299 {
		return build, fmt.Errorf(
			"response failed with status code: %d and body: %s",
			res.StatusCode,
			body,
		)
	}

	err = json.Unmarshal(body, &build)
	if err != nil {
		return build, err
	}

	return build, nil
}
