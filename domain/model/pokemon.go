package model

import "strconv"

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

func (this *Pokemon) ParseToPokemon(record []string) *Pokemon {
	id, _ := strconv.ParseInt(record[0], 10, 64)
	total, _ := strconv.ParseInt(record[3], 10, 32)
	hp, _ := strconv.ParseInt(record[4], 10, 32)
	attack, _ := strconv.ParseInt(record[5], 10, 32)
	defense, _ := strconv.ParseInt(record[6], 10, 32)
	spAtk, _ := strconv.ParseInt(record[7], 10, 32)
	spDef, _ := strconv.ParseInt(record[8], 10, 32)
	speed, _ := strconv.ParseInt(record[9], 10, 32)
	generation, _ := strconv.ParseInt(record[10], 10, 32)

	return &Pokemon{
		Id:         id,
		Name:       record[1],
		Type:       record[2],
		Total:      total,
		HP:         hp,
		Attack:     attack,
		Defense:    defense,
		SpAtk:      spAtk,
		SpDef:      spDef,
		Speed:      speed,
		Generation: generation,
	}

}
