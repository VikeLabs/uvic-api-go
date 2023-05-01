package schemas

type Building struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type BuildingSummary struct {
	Building string        `json:"building"`
	Data     []RoomSummary `json:"data"`
}

type RoomSummary struct {
	ID        uint64      `json:"room_id"`
	Room      string      `json:"room"`
	NextClass interface{} `json:"next_class"`
	Subject   interface{} `json:"subject"`
}
