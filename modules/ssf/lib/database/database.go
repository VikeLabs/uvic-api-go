package database

import (
	"context"
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	Sections  string = "sections"
	Buildings string = "buildings"
	Rooms     string = "rooms"
)

func New(ctx context.Context) *gorm.DB {
	path := getDSN()
	db, err := gorm.Open(sqlite.Open(path))
	if err != nil {
		panic(err)
	}
	return db.WithContext(ctx)
}

func getDSN() string {
	p := []string{"modules", "ssf", "database.db"}
	path := strings.Join(p, string(os.PathSeparator))
	return path
}
