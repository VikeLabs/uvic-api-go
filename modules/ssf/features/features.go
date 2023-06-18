package features

import (
	"github.com/VikeLabs/uvic-api-go/database"
	"gorm.io/gorm"
)

type state struct {
	*gorm.DB
}

func New() (*state, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	return &state{db}, nil
}
