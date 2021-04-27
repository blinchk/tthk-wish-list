package wishes

import (
	"net/http"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Receive(c *gin.Context) {
	var wishes []models.Wish
	var user models.User
	var err error
	c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	}
	err, wishes = database.GetWishes(user)
	message := gin.H{"success": true, "count": len(wishes)}
	c.JSON(http.StatusOK, message)
}
