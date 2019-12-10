package polo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	polo = "polo"
)

//Marco for Marco
func Marco(c *gin.Context) {
	c.String(http.StatusOK, polo)
}
