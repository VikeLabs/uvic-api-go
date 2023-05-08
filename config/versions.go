package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Versions struct {
	Database struct {
		Banner string
	}
}

func GetVersions() (*Versions, error) {
	var buf Versions

	f, err := os.ReadFile("./config/versions.yml")
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(f, &buf); err != nil {
		return nil, err
	}

	return &buf, nil
}
