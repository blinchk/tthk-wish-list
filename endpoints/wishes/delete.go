package wishes

import (
	"github.com/bredbrains/tthk-wish-list/models"
	"net/http"
	"strconv"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var id int
	var err error
	var wish models.Wish
	var allowed bool
	var currentUser models.User
	var gift models.Gift
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Bad ID of wish."})
		return
	}
	err, wish = database.GetWish(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	err, allowed = CheckWishPermissions(wish, c.GetHeader("Token"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	if !allowed {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"success": false, "error": "This action is not allowed for you."})
		return
	}
	err, currentUser = database.UserData(c.GetHeader("Token"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	err, gift = database.GetGiftByWish(wish, currentUser)
	if gift.ID != 0 {
		err = database.DeleteGift(gift)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
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
