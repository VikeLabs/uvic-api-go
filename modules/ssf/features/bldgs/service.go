package bldg

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/VikeLabs/uvic-api-go/lib/api"
	"github.com/VikeLabs/uvic-api-go/modules/ssf/schemas"
	"gorm.io/gorm"
)

func bldgsService() ([]schemas.Building, *api.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := newDB(ctx)

	var bldgs []schemas.Building
	if err := db.getBuildings(&bldgs); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, api.ErrNotFound(err, "Buildings not found")
		}
		return nil, api.ErrInternalServer(err)
	}

	for i := 0; i < len(bldgs); i++ {
		bldgs[i].Name = strings.ReplaceAll(bldgs[i].Name, "&amp;", "&")
	}

	return bldgs, nil
}
