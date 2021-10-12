package services

import (
	"strconv"

	. "bootcamp/domain/model"
)

//Parses an array of strings to a pokemon pointer
func ParseToPokemonPointer(row []string) *Pokemon {
	id, _ := strconv.ParseInt(row[0], 10, 64)
	total, _ := strconv.ParseInt(row[3], 10, 32)
	hp, _ := strconv.ParseInt(row[4], 10, 32)
	attack, _ := strconv.ParseInt(row[5], 10, 32)
	defense, _ := strconv.ParseInt(row[6], 10, 32)
	spAtk, _ := strconv.ParseInt(row[7], 10, 32)
	spDef, _ := strconv.ParseInt(row[8], 10, 32)
	speed, _ := strconv.ParseInt(row[9], 10, 32)
	generation, _ := strconv.ParseInt(row[10], 10, 32)
	return &Pokemon{
		Id:         id,
		Name:       row[1],
		Type:       row[2],
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

//Parses an array of strings to a pokemon struct
func ParseToPokemon(row []string) Pokemon {
	id, _ := strconv.ParseInt(row[0], 10, 64)
	total, _ := strconv.ParseInt(row[3], 10, 32)
	hp, _ := strconv.ParseInt(row[4], 10, 32)
	attack, _ := strconv.ParseInt(row[5], 10, 32)
	defense, _ := strconv.ParseInt(row[6], 10, 32)
	spAtk, _ := strconv.ParseInt(row[7], 10, 32)
	spDef, _ := strconv.ParseInt(row[8], 10, 32)
	speed, _ := strconv.ParseInt(row[9], 10, 32)
	generation, _ := strconv.ParseInt(row[10], 10, 32)
	return Pokemon{
		Id:         id,
		Name:       row[1],
		Type:       row[2],
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
