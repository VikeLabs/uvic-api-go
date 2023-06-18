package database

import (
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	databaseFile string
)

func init() {
	databaseFile = strings.Join(
		[]string{"database", "database.db"},
		string(os.PathSeparator),
	)
}

const (
	Sections  string = "sections"
	Buildings string = "buildings"
	Rooms     string = "rooms"
)

func New() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(databaseFile))
	if err != nil {
		return nil, err
	}
	return db, nil
}
