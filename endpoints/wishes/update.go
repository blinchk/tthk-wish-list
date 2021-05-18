package wishes

import (
	"net/http"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var wish models.Wish
	var err error
	err = c.BindJSON(&wish)
	var allowed bool
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Request body is invalid."})
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
	err, wish = database.UpdateWish(wish)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	message := gin.H{"wish": wish}
	c.JSON(http.StatusOK, message)
	return
}
