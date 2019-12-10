package polo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMarco(t *testing.T) {

	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)
	req, _ := http.NewRequest(http.MethodGet, "/marco", nil)
	c.Request = req
	Marco(c)
	assert.EqualValues(t, http.StatusOK, res.Code)
	assert.EqualValues(t, polo, res.Body.String())
}
