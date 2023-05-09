package database

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	Sections  string = "sections"
	Buildings string = "buildings"
	Rooms     string = "rooms"
)

func New(ctx context.Context) *gorm.DB {
	path := "database.db"
	db, err := gorm.Open(sqlite.Open(path))
	if err != nil {
		panic(err)
	}
	return db.WithContext(ctx)
}
