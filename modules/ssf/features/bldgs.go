package features

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/VikeLabs/uvic-api-go/lib/api"
	"gorm.io/gorm"
)

func BldgsController(w http.ResponseWriter, r *http.Request) {
	bldgs, err := bldgsService()
	if err != nil {
		err.HandleError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&bldgs)
}

func bldgsService() ([]Building, *api.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := newDB(ctx)

	var bldgs []Building
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
