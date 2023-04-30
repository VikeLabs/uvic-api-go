package bldgid

import (
	"context"
	"errors"

	"github.com/VikeLabs/uvic-api-go/modules/ssf/lib"
	"github.com/VikeLabs/uvic-api-go/modules/ssf/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	ErrBadQuery error = errors.New("bad query param")
)

type model struct {
	*gorm.DB
}

type roomId struct {
	RoomID int
}

func (db *model) getRooms(bldg *schemas.Building, rooms *schemas.RoomSummary) error {
	return nil
}

func (db *model) getBuildingName(bldg *schemas.Building) error {
	if bldg.ID == 0 {
		return ErrBadQuery
	}
	return db.Where("id=?", bldg.ID).First(bldg).Error
}

func newDB(ctx context.Context) *model {
	path := lib.GetDSN()
	db, err := gorm.Open(sqlite.Open(path))
	if err != nil {
		panic(err)
	}
	return &model{db.WithContext(ctx)}
}
