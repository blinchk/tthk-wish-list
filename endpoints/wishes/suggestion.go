package wishes

import (
	"net/http"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Suggestion(c *gin.Context) {
	var wishes []models.Wish
	var follow models.Follow
	var err error
	c.BindJSON(&follow)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	}
	err, wishes = database.GetSuggestion(follow)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	}
	message := gin.H{"success": true, "wish 1": wishes[0].ID, "wish 2": wishes[1].ID, "wish 3": wishes[2].ID}
	c.JSON(http.StatusOK, message)
}
