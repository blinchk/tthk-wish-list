package wishes

import (
	"net/http"

	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
	var wish models.Wish
	var err error
	c.BindJSON(&wish)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	}
	message := gin.H{"success": true}
	c.JSON(http.StatusOK, message)
}
