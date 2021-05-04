package wishes

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var wish models.Wish
	var err error
	err = c.BindJSON(&wish)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	v := validator.New()
	err = v.Struct(wish)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err, wish = database.AddWish(wish)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	message := gin.H{"success": true, "wish": wish}
	c.JSON(http.StatusOK, message)
	return
}
