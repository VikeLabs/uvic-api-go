package ssf

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Session struct {
	ID           uint64 `json:"id"`
	Subject      string `json:"subject"`
	Description  string `json:"description"`
	TimeStartStr string `json:"time_start"`
	TimeEndStr   string `json:"time_end"`
	Room         string `json:"-" gorm:"-"`
	TimeStartInt uint64 `json:"-"`
	TimeEndInt   uint64 `json:"-"`
}

type Room struct {
	ID           uint64   `json:"id"`
	Room         string   `json:"room"`
	NextClass    *Session `json:"next_class" gorm:"-"`
	CurrentClass *Session `json:"current_class" gorm:"-"`
}

type TimeQuery struct {
	hour   int
	minute int
	day    int
}

func routeBuildingID(c *fiber.Ctx) error {
	buildingID, err := strconv.ParseInt(c.Params("buildingID"), 10, 32)

	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	log.Println(c.Queries())
	log.Println(buildingID)

	q, err := parseQuery(c)
	log.Println(q)

	return nil
}

func parseQuery(c *fiber.Ctx) (*TimeQuery, error) {
	q := c.Queries()
	h, ok := q["hour"]
	if !ok {
		return nil, errors.New("missing hour query")
	}
	hour, err := strconv.ParseInt(h, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse hour: %v", err)
	}

	m, ok := q["minute"]
	if !ok {
		return nil, errors.New("missing minute query")
	}
	minute, err := strconv.ParseInt(m, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse minute: %v", err)
	}

	day, err := strconv.ParseInt(q["day"], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse day: %v", err)
	}

	return &TimeQuery{hour: int(hour), minute: int(minute), day: int(day)}, nil
}
