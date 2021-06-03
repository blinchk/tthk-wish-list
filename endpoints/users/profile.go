package users

import (
	"net/http"
	"strconv"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/bredbrains/tthk-wish-list/modules"
	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	var wishes []models.Wish
	var user models.User
	var requestingUser models.User
	var err error
	var following, isSameUser bool
	id, err := strconv.Atoi(c.Param("id"))
	token := c.GetHeader("Token")
	if token != "" {
		err, requestingUser = database.UserData(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid token"})
			return
		}
		following, isSameUser = modules.CheckIsFollowed(requestingUser, id)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	err, user = database.UserDataById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Requested user isn't exists."})
		return
	}
	err, wishes = database.GetWishes(user, requestingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	if token != "" && !isSameUser {
		c.JSON(http.StatusOK, gin.H{"success": true, "user": user, "wishes": wishes, "following": following})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "user": user, "wishes": wishes})
	}
	return
}

func EditUserProfile(c *gin.Context) {
	var wishes []models.Wish
	var newUser models.User
	var user models.User
	var err error
	err = c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request body."})
	}
	accessToken := c.GetHeader("Token")
	err, user = database.UserData(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "You don't have permissions for this."})
		return
	}
	if len(newUser.FirstName) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "This fields cannot be blank."})
		return
	}
	newUser.ID = user.ID
	err, user = database.EditUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Requested user isn't exists."})
		return
	}
	err, wishes = database.GetWishes(user, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "user": user, "wishes": wishes})
	return
}
