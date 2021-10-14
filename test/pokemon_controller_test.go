package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	. "bootcamp/domain/model"
	. "bootcamp/interface/controller"
)

type mockPkmnResponse struct {
	Data   []Pokemon `json:"data"`
	Length int       `json:"length"`
}
type mockPokemonController struct{}

var mockPkmn = Pokemon{Id: 1, Name: "Bulbasaur", Type: "Grass", Total: 318, HP: 45, Attack: 49, Defense: 49, SpAtk: 65, SpDef: 65, Speed: 45, Generation: 1}

//Get pokemon mock function
func (m *mockPokemonController) Get() gin.HandlerFunc {
	gin.SetMode(gin.TestMode)

	var response = []Pokemon{mockPkmn}
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, response)
	}
}

//Get async pokemon mock function
func (m *mockPokemonController) GetAsync() gin.HandlerFunc {
	gin.SetMode(gin.TestMode)

	var response = []Pokemon{mockPkmn}
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": response, "length": len(response)})
	}
}

//Mock Controller instance
func newMockPokemonController(t *testing.T) *mockPokemonController {
	return &mockPokemonController{}
}

var ctrller PokemonController

//Run test
func TestGetPokemonContoller(t *testing.T) {
	ctrller = newMockPokemonController(t)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.GET("/pokemon/get", ctrller.Get())

	c.Request, _ = http.NewRequest(http.MethodGet, "/pokemon/get", nil)
	r.ServeHTTP(w, c.Request)

	data := []Pokemon{}
	json.Unmarshal(w.Body.Bytes(), &data)

	assert.EqualValues(t, w.Code, http.StatusOK)
	assert.EqualValues(t, data, []Pokemon{mockPkmn})
}

//Run test async
func TestGetAsyncPokemonContoller(t *testing.T) {
	ctrller = newMockPokemonController(t)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.GET("/pokemon/get/async", ctrller.GetAsync())

	c.Request, _ = http.NewRequest(http.MethodGet, "/pokemon/get/async?type=odd&items=1&items_per_workers=1", nil)
	r.ServeHTTP(w, c.Request)

	data := mockPkmnResponse{}
	json.Unmarshal(w.Body.Bytes(), &data)

	assert.EqualValues(t, w.Code, http.StatusOK)
	assert.EqualValues(t, data.Length, 1)
}
