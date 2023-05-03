package schemas

type Building struct {
	ID   uint64 `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
}

type BuildingSummary struct {
	Building string        `json:"building"`
	Data     []RoomSummary `json:"data"`
}

type RoomSummary struct {
	ID        uint64  `json:"room_id"`
	Room      string  `json:"room"`
	NextClass *string `json:"next_class"`
	Subject   *string `json:"subject"`
}

type Room struct {
	ID         uint64   `json:"id"`
	Room       string   `json:"room"`
	Building   Building `json:"-"`
	BuildingID uint64   `json:"building_id"`
}
