package wishes

import (
	"net/http"
	"strconv"

	"github.com/bredbrains/tthk-wish-list/models"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var err error
	var user models.User
	var wish models.Wish
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Bad ID of wish."})
		return
	}
	err, wish = database.GetWish(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
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
	err = database.DeleteWish(wish)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	message := gin.H{"success": true}
	c.JSON(http.StatusOK, message)
	return
}
