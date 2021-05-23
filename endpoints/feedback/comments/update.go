package comments

import (
	"net/http"
	"strconv"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var comment models.Comment
	var mutable models.Comment
	var err error
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Bad ID of comment."})
		return
	}
	token := c.GetHeader("Token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	err, user = database.UserData(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	err, comment = database.GetComment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	if comment.User.ID != user.ID {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "No permission."})
		return
	}
	err = c.BindJSON(&mutable)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Request body is invalid."})
		return
	}
	comment.Content = mutable.Content
	err, comment = database.UpdateComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	message := gin.H{"success": true, "comment": comment}
	c.JSON(http.StatusOK, message)
	return
}
