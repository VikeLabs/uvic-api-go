package database

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbVersion struct {
	StudySpaceFinder string `yaml:"ssf"`
}

var (
	versionFile  string
	databaseFile string
)

func init() {
	databaseFile = strings.Join(
		[]string{"database", "database.db"},
		string(os.PathSeparator),
	)

	versionFile = strings.Join(
		[]string{"database", "versions.yml"},
		string(os.PathSeparator),
	)
}

const (
	Sections  string = "sections"
	Buildings string = "buildings"
	Rooms     string = "rooms"
)

func New(ctx context.Context) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(databaseFile))
	if err != nil {
		log.Fatal(err)
	}
	return db.WithContext(ctx)
}

func GetVersion() (*DbVersion, error) {
	path := strings.Join(
		[]string{"database", "versions.yml"},
		string(os.PathSeparator),
	)

	f, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
		return nil, err
	}

	var v DbVersion
	if err := yaml.Unmarshal(f, &v); err != nil {
		return nil, err
	}

	return &v, nil
}

func SetVersion(v DbVersion) error {
	buf, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	msg := []byte("# This file is managed via scripts, don't edit\n\n")

	msg = append(msg, buf...)
	if err := os.WriteFile(versionFile, msg, 0666); err != nil {
		panic(err)
	}

	return nil
}
