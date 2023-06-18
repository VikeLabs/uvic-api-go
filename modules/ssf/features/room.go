package features

import (
	"net/http"

	"github.com/VikeLabs/uvic-api-go/lib/api"
	"github.com/VikeLabs/uvic-api-go/modules/ssf/schemas"
	"github.com/go-chi/chi/v5"
)

type roomSchedules map[string][]scheduleEntry

type scheduleEntry struct {
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
	Subject   string `json:"subject"`
}

type scheduleQuery struct {
	schemas.Section
	Subject string
}

func (db *state) GetRoomSchedule(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "id")
	if roomID == "" {
		err := api.ErrBadRequest(nil, "missing url param: room id")
		err.HandleError(w)
		return
	}

	var data []scheduleQuery
	if err := db.getRoomSchedules(roomID, &data); err != nil {
		api.ErrInternalServer(err).HandleError(w)
		return
	}

	response := roomSchedules{}
	for _, v := range data {
		response.sortSection(v)
	}

	api.ResponseBuilder(w).
		Status(http.StatusOK).
		JSON(response)
}

func (db *state) getRoomSchedules(roomID string, data *[]scheduleQuery) error {
	sel := []string{
		"sections.time_start_str",
		"sections.time_end_str",
		"sections.monday",
		"sections.tuesday",
		"sections.wednesday",
		"sections.thursday",
		"sections.friday",
		"sections.saturday",
		"sections.sunday",
		"subjects.subject",
	}
	db.Table(schemas.TableSection).
		Select(sel).
		Where("room_id = ?", roomID).
		InnerJoins("INNER JOIN subjects ON subjects.id=sections.subject_id").
		Find(data)

	return nil
}

func (r *roomSchedules) sortSection(s scheduleQuery) {
	detail := scheduleEntry{
		TimeStart: s.TimeEndStr,
		TimeEnd:   s.TimeEndStr,
		Subject:   s.Subject,
	}

	if s.Monday {
		(*r)["Monday"] = append((*r)["Monday"], detail)
	}

	if s.Tuesday {
		(*r)["Tuesday"] = append((*r)["Tuesday"], detail)
	}

	if s.Wednesday {
		(*r)["Wednesday"] = append((*r)["Wednesday"], detail)
	}

	if s.Thursday {
		(*r)["Thursday"] = append((*r)["Thursday"], detail)
	}

	if s.Friday {
		(*r)["Friday"] = append((*r)["Friday"], detail)
	}

	if s.Saturday {
		(*r)["Saturday"] = append((*r)["Saturday"], detail)
	}

	if s.Sunday {
		(*r)["Sunday"] = append((*r)["Sunday"], detail)
	}
}
