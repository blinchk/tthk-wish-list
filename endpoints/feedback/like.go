package feedback

import (
	"net/http"
	"strconv"
	"time"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func GetLike(c *gin.Context) {
	var like models.Like
	var message gin.H
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Bad ID of like."})
		return
	}
	err, like = database.GetLike(id)
	if err != nil {
		message = gin.H{"success": true, "liked": false}
	} else {
		message = gin.H{"success": true, "liked": true, "like": like}
	}
	c.JSON(http.StatusOK, message)
	return
}

func GetLikesCount(c *gin.Context) {
	var like models.Like
	var count int
	var message gin.H
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Bad ID of like."})
		return
	}
	conType := c.Param("type")
	like.Connection, like.ConnectionType = id, conType
	err, count = database.GetLikesCount(like)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	message = gin.H{"success": true, "count": count}
	c.JSON(http.StatusOK, message)
	return
}

func ToggleLike(c *gin.Context) {
	var like models.Like
	var liked bool
	var err error
	var message gin.H
	var count int
	err = c.BindJSON(&like)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Request body is invalid."})
		return
	}
	token := c.GetHeader("Token")
	err, like.User = database.UserData(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	if database.LikeExist(like) {
		err = database.DeleteLike(like)
		liked = false
		err, count = database.GetLikesCount(like)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		message = gin.H{"success": true, "liked": liked, "count": count}
	} else {
		like.CreationTime = time.Now().Format("2006-01-02 15:04:05")
		err, like = database.AddLike(like)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		liked = true
		err, count = database.GetLikesCount(like)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		message = gin.H{"success": true, "like": like, "liked": liked, "count": count}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, message)
	return
}
