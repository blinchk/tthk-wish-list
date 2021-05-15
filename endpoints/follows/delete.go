package follows

import (
	"net/http"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var follow models.Follow
	var err error
	err = c.BindJSON(&follow)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Request body is invalid."})
		return
	}
	err, follow = database.DeleteFollow(follow)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	message := gin.H{"follow": follow}
	c.JSON(http.StatusOK, message)
	return
}
