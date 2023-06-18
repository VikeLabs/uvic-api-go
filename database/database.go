package database

import (
	"context"
	"log"
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

func New(ctx context.Context) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(databaseFile))
	if err != nil {
		log.Fatal(err)
	}
	return db.WithContext(ctx)
}
