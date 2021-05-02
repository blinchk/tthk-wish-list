package users

import (
	"fmt"
	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Wishes(c *gin.Context) {
	var wishes []models.Wish
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	user := models.User{ID: id}
	err, wishes = database.GetWishes(user)
	err, user = database.UserDataById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	message := gin.H{"success": true, "wishes": wishes, "user": user}
	c.JSON(http.StatusOK, message)
	return
}
