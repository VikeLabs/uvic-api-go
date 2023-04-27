package features

import (
	"context"
	"strings"
	"time"
)

func GetBuildings() ([]Building, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := newDB(ctx)

	var bldgs []Building
	if err := db.getBuildings(&bldgs); err != nil {
		return nil, err
	}

	for i := 0; i < len(bldgs); i++ {
		bldgs[i].Name = strings.ReplaceAll(bldgs[i].Name, "&amp;", "&")
	}

	return bldgs, nil
}
