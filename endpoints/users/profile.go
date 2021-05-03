package users

import (
	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

func GetUserProfile(c *gin.Context) {
	var wishes []models.Wish
	var user models.User
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	err, user = database.UserDataById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Requested user isn't exists."})
		return
	}
	err, wishes = database.GetWishes(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "user": user, "wishes": wishes})
	return
}

func EditUserProfile(c *gin.Context) {
	var wishes []models.Wish
	var newUser models.User
	c.BindJSON(&newUser)
	accessToken := c.GetHeader("Token")
	err, user := database.UserData(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "You don't have permissions for this."})
		return
	}
	v := validator.New()
	err = v.Struct(user)
	newUser.ID = user.ID
	err, user = database.EditUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Requested user isn't exists."})
		return
	}
	err, wishes = database.GetWishes(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "user": user, "wishes": wishes})
	return
}
