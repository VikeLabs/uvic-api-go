package bldg

import (
	"context"
	"errors"

	"github.com/VikeLabs/uvic-api-go/modules/ssf/lib/database"
	"github.com/VikeLabs/uvic-api-go/modules/ssf/schemas"
	"gorm.io/gorm"
)

type model struct {
	*gorm.DB
}

func newDB(ctx context.Context) *model {
	db := database.New(ctx)
	return &model{db.WithContext(ctx)}
}

func (db *model) getBuildings(bldgs *[]schemas.Building) error {
	trx := db.Table(database.Buildings)
	trx.Order("name ASC")
	trx.Find(bldgs)

	if err := trx.Error; err != nil {
		return err
	}

	if trx.RowsAffected == 0 {
		return errors.New("not found") // TODO: change this to return api error
	}
	return nil
}
