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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Bad ID of user."})
		return
	}
	gift.User.ID = id
	err = c.BindJSON(&gift)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Request body is invalid."})
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
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Bad ID of wish."})
		return
	}
	err, gifts = database.GetGiftsByWish(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "gifts": gifts})
	return
}
