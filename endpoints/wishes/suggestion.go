package wishes

import (
	"math/rand"
	"net/http"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Suggestion(c *gin.Context) {
	var wishes []models.Wish
	var user models.User
	var err error
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Request body is invalid."})
		return
	}
	min := 0
	max := len(database.GetFollowsFromUser(user))
	rndInt := rand.Intn(max-min) + min
	err, wishes = database.GetSuggestion(database.GetFollowsFromUser(user)[rndInt])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	message := gin.H{"wishes": wishes}
	c.JSON(http.StatusOK, message)
	return
}
