package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	. "bootcamp/infraestructure/router"
	. "bootcamp/interface/controller"
)

func TestPing(t *testing.T) {
	var mockCatctrller CatController
	mockCatctrller = newMockCatController(t)

	var mockPkmnCtrller PokemonController
	mockPkmnCtrller = newMockPokemonController(t)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	router := SetupRouter(r, mockCatctrller, mockPkmnCtrller)

	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
