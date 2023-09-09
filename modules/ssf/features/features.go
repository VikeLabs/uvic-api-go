package features

import (
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type state struct {
	*gorm.DB
}

const (
	tableSections  string = "sections"
	tableBuildings string = "buildings"
	tableRooms     string = "rooms"
)

var (
	databaseFile string
)

func init() {
	databaseFile = strings.Join(
		[]string{"modules", "ssf", "database.db"},
		string(os.PathSeparator),
	)
}

func New() (*state, error) {
	db, err := gorm.Open(sqlite.Open(databaseFile))
	if err != nil {
		return nil, err
	}

	return &state{db}, nil
}
