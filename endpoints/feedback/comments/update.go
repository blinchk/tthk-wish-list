package comments

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	message := gin.H{"success": true}
	c.JSON(http.StatusOK, message)
	return
}
