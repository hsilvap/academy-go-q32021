package model

type Pokemon struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Total      int64  `json:"total"`
	HP         int64  `json:"hp"`
	Attack     int64  `json:"attack"`
	Defense    int64  `json:"defense"`
	SpAtk      int64  `json:"spAtk"`
	SpDef      int64  `json:"spDef"`
	Speed      int64  `json:"speed"`
	Generation int64  `json:"generation"`
}
