package feedback

import (
	"net/http"
	"strconv"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func ToggleGift(c *gin.Context) {
	var gift models.Gift
	var err error
	var message gin.H
	user, err := strconv.Atoi(c.Param("user"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Bad ID of user."})
		return
	}
	wish, err := strconv.Atoi(c.Param("wish"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Bad ID of user."})
		return
	}
	err = c.BindJSON(&gift)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Request body is invalid."})
		return
	}
	err, currentUser := database.UserData(c.GetHeader("Token"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid token."})
		return
	}
	err, gift.User = database.UserDataById(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	err, gift.Wish = database.GetWishByIdAndUser(wish, currentUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	if database.GiftExist(gift) {
		err = database.DeleteGift(gift)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		message = gin.H{"success": true}
	} else {
		err, gift = database.AddGift(gift)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		message = gin.H{"success": true, "gift": gift}
	}
	c.JSON(http.StatusOK, message)
	return
}

func GetGifts(c *gin.Context) {
	var gifts []models.Gift
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Bad ID of user."})
		return
	}
	err, currentUser := database.UserData(c.GetHeader("Token"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid token."})
		return
	}
	err, gifts = database.GetGiftsByUsers(id, currentUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "gifts": gifts})
	return
}
