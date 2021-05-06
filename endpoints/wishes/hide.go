package wishes

import (
	"net/http"
	"strconv"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Hide(c *gin.Context) {
	var wish models.Wish
	var err error
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Bad ID of wish."})
		return
	}
	err, wish = database.GetWish(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err, user = database.UserData(c.GetHeader("Token"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	if wish.User.ID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "You don't have permissions for this."})
		return
	}
	err, wish = database.HideWish(wish)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	message := gin.H{"success": true, "wish": wish}
	c.JSON(http.StatusOK, message)
	return
}
