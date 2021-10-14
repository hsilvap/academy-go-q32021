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

type mockCatController struct{}

var mockCat = Cat{Id: "15t", Url: "https://cdn2.thecatapi.com/images/14t.gif", Width: 246, Height: 233}

//Get cat mock function
func (m *mockCatController) Get() gin.HandlerFunc {
	gin.SetMode(gin.TestMode)

	var response = []Cat{mockCat}
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, response)
	}
}

//Mock Controller instance
func newMockCatController(t *testing.T) *mockCatController {
	return &mockCatController{}
}

//Run test
func TestCatContoller(t *testing.T) {
	var ctrller CatController
	ctrller = newMockCatController(t)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.GET("/cat", ctrller.Get())

	c.Request, _ = http.NewRequest(http.MethodGet, "/cat", nil)
	r.ServeHTTP(w, c.Request)

	data := []Cat{}
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.EqualValues(t, w.Code, http.StatusOK)
	assert.EqualValues(t, data, []Cat{mockCat})
}
