package bldg

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/VikeLabs/uvic-api-go/modules/ssf/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type database struct {
	*gorm.DB
}

func newDB(ctx context.Context) *database {
	p := []string{"modules", "ssf", "database.db"}
	path := strings.Join(p, string(os.PathSeparator))

	db, err := gorm.Open(sqlite.Open(path))
	if err != nil {
		panic(err)
	}

	return &database{db.WithContext(ctx)}
}

func (db *database) getBuildings(bldgs *[]schemas.Building) error {
	result := db.Order("name ASC").Find(bldgs)
	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return errors.New("not found") // TODO: change this to return api error
	}

	return nil
}
