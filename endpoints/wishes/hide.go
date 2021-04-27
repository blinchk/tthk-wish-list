package wishes

import (
	"net/http"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Hide(c *gin.Context) {
	var wish models.Wish
	var err error
	c.BindJSON(&wish)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	}
	err = database.HideWish(wish)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	}
	message := gin.H{"success": true}
	c.JSON(http.StatusOK, message)
}
