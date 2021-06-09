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
	wish, err := strconv.Atoi(c.Param("wish"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Bad ID of wish."})
		return
	}
	err, currentUser := database.UserData(c.GetHeader("Token"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid token."})
		return
	}
	err, gift.Wish = database.GetWishByIdAndUser(wish, currentUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	gift.User = currentUser
	gift.Wish.ID = wish
	if database.GiftExist(gift) {
		err = database.DeleteGift(gift)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		message = gin.H{"success": true}
	} else {
		err = c.BindJSON(&gift)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Request body is invalid."})
			return
		}
		err, gift = database.AddGift(gift)
		gift.User.Email = ""
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		message = gin.H{"success": true, "gift": gift}
	}
	c.JSON(http.StatusOK, message)
	return
}

func GetGiftsByUser(c *gin.Context) {
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

func GetGifts(c *gin.Context) {
	var gifts []models.Gift
	err, currentUser := database.UserData(c.GetHeader("Token"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid token."})
		return
	}
	err, gifts = database.GetGifts(currentUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "gifts": gifts})
	return
}

func EditGift(c *gin.Context) {
	var gift models.Gift
	var err error
	var currentUser models.User
	var currentGift models.Gift
	err = c.BindJSON(&gift)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid gift handled."})
	}
	err, currentUser = database.UserData(c.GetHeader("Token"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid token."})
		return
	}
	gift.User = currentUser
	err, currentGift = database.GetGift(gift)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Gift does not exist."})
		return
	}
	if currentGift.User.ID != currentUser.ID {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"success": false, "error": "Not allowed for you."})
		return
	} else {
		err, gift = database.EditGift(gift)
		c.JSON(http.StatusOK, gin.H{"success": true, "gift": gift})
	}
}
