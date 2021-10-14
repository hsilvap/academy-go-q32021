package model

type PokemonAsyncUriQueryParams struct {
	Type           string `form:"type" binding:"required"`
	Items          int    `form:"items" binding:"required"`
	ItemsPerWorker int    `form:"items_per_workers" binding:"required"`
}
