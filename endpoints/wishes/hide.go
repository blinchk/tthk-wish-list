package wishes

import (
	"net/http"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Hide(c *gin.Context) {
	var pool models.Pool
	var err error
	c.BindJSON(&pool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	}
	err = database.HideWish(pool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	}
	message := gin.H{"success": true}
	c.JSON(http.StatusOK, message)
}
