package wishes

import (
	"github.com/go-playground/validator/v10"
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
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Request body is invalid."})
		return
	}
	v := validator.New()
	err = v.Struct(wish)
	if wish.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Wish name cannot be blank."})
		return
	}
	err, wish = AssignUserToWish(wish, c.GetHeader("Token"))
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"success": false, "error": "This action is not allowed for you."})
		return
	}
	err, wish = database.AddWish(wish)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid wish."})
		return
	}
	message := gin.H{"success": true, "wish": wish}
	c.JSON(http.StatusOK, message)
	return
}
