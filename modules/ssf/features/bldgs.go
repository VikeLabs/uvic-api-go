package features

import (
	"net/http"
	"strings"

	"github.com/VikeLabs/uvic-api-go/database"
	"github.com/VikeLabs/uvic-api-go/lib/api"
	"github.com/VikeLabs/uvic-api-go/modules/ssf/schemas"
	"gorm.io/gorm"
)

type buildingQueries struct {
	schemas.Building
}

// replace building name interface for gorm
func (b *buildingQueries) AfterFind(tx *gorm.DB) error {
	b.Name = strings.ReplaceAll(b.Name, "&amp;", "&")
	return nil
}

func (db *state) Buildings(w http.ResponseWriter, r *http.Request) {
	var bldgs []buildingQueries
	if err := db.queryBuildings(&bldgs); err != nil {
		api.ResponseBuilder(w).
			Error(api.ErrInternalServer(err))
		return
	}

	api.ResponseBuilder(w).Status(http.StatusOK).JSON(bldgs)
}

func (db *state) queryBuildings(buf *[]buildingQueries) error {
	return db.Table(database.Buildings).
		Find(buf).Error
}
