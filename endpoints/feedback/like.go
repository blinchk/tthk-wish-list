package feedback

import (
	"net/http"
	"time"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func GetLike(c *gin.Context) {

}

func ToggleLike(c *gin.Context) {
	var like models.Like
	var err error
	var message gin.H
	err = c.BindJSON(&like)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Request body is invalid."})
		return
	}
	token := c.GetHeader("Token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	err, like.User = database.UserData(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	if database.LikeExist(like) {
		database.DeleteLike(like)
		message = gin.H{"success": true, "liked": false}
	} else {
		like.CreationTime = time.Now().Format("2006-01-02 15:04:05")
		database.AddLike(like)
		message = gin.H{"success": true, "liked": true, "like": like}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, message)
}
