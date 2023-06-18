package features

import (
	"context"

	"github.com/VikeLabs/uvic-api-go/database"
	"gorm.io/gorm"
)

type state struct {
	*gorm.DB
}

var handlers state

func New(ctx context.Context) state {
	db := database.New(ctx)
	handlers = state{db}
	return handlers
}
