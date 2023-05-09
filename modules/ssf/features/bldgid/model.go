package bldgid

import (
	"context"
	"errors"

	"github.com/VikeLabs/uvic-api-go/lib/database"
	"github.com/VikeLabs/uvic-api-go/modules/ssf/lib"
	"github.com/VikeLabs/uvic-api-go/modules/ssf/schemas"
	"gorm.io/gorm"
)

var (
	ErrNoData   error = errors.New("no data")
	ErrBadQuery error = errors.New("bad query param")
)

type model struct {
	*gorm.DB
}

type Session struct {
	TimeStartStr string `json:"time_start_str"`
	ID           uint64 `json:"room_id"`
	Room         string `json:"room" gorm:"room"`
	Subject      string `json:"subject"`
	TimeStartInt uint64 `json:"-"`
	TimeEndInt   uint64 `json:"-"`
}

func (db *model) getRoomSchedule(roomID uint64, query *lib.TimeQueries, buf *[]Session) error {
	sel := []string{
		"sections.time_start_str",
		"rooms.id",
		"rooms.room",
		"subjects.subject",
		"sections.time_start_int",
		"sections.time_end_int",
	}

	filter := map[string]any{
		"sections.room_id": roomID,
		query.Day:          true,
	}

	sql := db.Table(database.Sections)
	sql.Select(sel)
	sql.Joins("JOIN rooms ON sections.room_id=rooms.id")
	sql.Joins("JOIN subjects ON sections.subject_id=subjects.id")
	sql.Where(filter)
	sql.Where("sections.time_start_int >= ?", query.Time)
	sql.Order("time_start_int ASC")
	sql.Limit(2)
	sql.Find(buf)

	if sql.RowsAffected == 0 {
		return ErrNoData
	}

	return sql.Error
}

func (db *model) getRooms(bldgID uint64, rooms *[]schemas.RoomSummary) error {
	sql := db.Table(database.Rooms)
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
	return &model{database.New(ctx)}
}
