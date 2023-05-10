package bldg

import (
	"context"
	"errors"

	"github.com/VikeLabs/uvic-api-go/database"
	"github.com/VikeLabs/uvic-api-go/modules/ssf/schemas"
	"gorm.io/gorm"
)

type model struct {
	*gorm.DB
}

func newDB(ctx context.Context) *model {
	return &model{database.New(ctx)}
}

func (db *model) getBuildings(bldgs *[]schemas.Building) error {
	result := db.Order("name ASC").Find(bldgs)
	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return errors.New("not found") // TODO: change this to return api error
	}

	return nil
}
