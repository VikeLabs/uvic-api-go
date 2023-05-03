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
	ErrNoData   error = errors.New("no data")
	ErrBadQuery error = errors.New("bad query param")
)

type model struct {
	*gorm.DB
}

type RoomSchedule struct {
	TimeStartStr string `json:"time_start_str"`
	RoomID       uint64 `json:"room_id"`
	RoomName     string `json:"room"`
	Subject      string `json:"subject"`
	TimeStartInt string `json:"-"`
	TimeEndInt   string `json:"-"`
}

func (db *model) getRoomSchedule(roomID uint64, day string, buf *RoomSchedule) error {
	sel := []string{
		"sections.time_start_str",
		"rooms.id",
		"rooms.room",
		"subjects.subject",
		"sections.time_start_int",
		"sections.time_end_int",
	}
	where := map[string]any{"sections.room_id": roomID, day: true} // TODO: change day

	sql := db.Table("sections")
	sql.Select(sel)
	sql.Joins("JOIN rooms ON sections.room_id=rooms.id")
	sql.Joins("JOIN subjects ON sections.subject_id=subjects.id")
	sql.Where(where)
	sql.Order("time_start_int ASC")
	sql.Find(buf)

	if sql.RowsAffected == 0 {
		return ErrNoData
	}

	return sql.Error
}

func (db *model) getRooms(bldgID uint64, rooms *[]schemas.RoomSummary) error {
	sql := db.Table("rooms")
	sql.Select("rooms.id", "rooms.room")
	sql.Joins("JOIN buildings ON rooms.building_id=buildings.id")
	sql.Where("buildings.id=?", bldgID)
	sql.Order("room ASC")
	sql.Scan(rooms)
	return sql.Error
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
