package database

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(ctx context.Context) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"))
	if err != nil {
		panic(err.Error())
	}

	return db.WithContext(ctx)
}
