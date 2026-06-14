package client

import (
	"encoding/json"
	"os"

	"github.com/theashishprasad/jenkins-build-report-tool/models"
)

// LoadBuild reads and parses build information from a JSON file.
func LoadBuild() (models.Build, error) {
	var build models.Build
	jsonData, err := os.ReadFile("sample/build.json")

	if err != nil {
		return build, err
	}

	err = json.Unmarshal(jsonData, &build)
	if err != nil {
		return build, err
	}

	return build, nil
}
