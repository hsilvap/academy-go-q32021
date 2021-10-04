package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	. "bootcamp/interface/controller"
)

type mockCatController struct{}

func (m *mockCatController) GetCat() gin.HandlerFunc {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return func(ctx *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "this worked"})
	}
}

//Mock Controller instance
func newMockCatController(t *testing.T) *mockCatController {
	return &mockCatController{}
}

// mock controller interface
func mockContoller(c CatController) {
	c.GetCat()
}

//Run test
func TestCatContoller(t *testing.T) {
	var ctrller CatController
	ctrller = newMockCatController(t)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.GET("/cat", ctrller.GetCat())

	c.Request, _ = http.NewRequest(http.MethodGet, "/cat", bytes.NewBuffer([]byte("{}")))
	r.ServeHTTP(w, c.Request)

	assert.EqualValues(t, w.Code, http.StatusOK)
}
