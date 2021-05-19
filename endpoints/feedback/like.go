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
	var mimic models.Like
	var liked bool
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
	mimic = like
	if database.LikeExist(like) {
		err = database.DeleteLike(like)
		liked = false
		message = gin.H{"success": true, "liked": false}
	} else {
		like.CreationTime = time.Now().Format("2006-01-02 15:04:05")
		err, like = database.AddLike(like)
		liked = true
		message = gin.H{"success": true, "like": like, "liked": true}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	err = database.UniteLike(mimic, liked)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, message)
	return
}
