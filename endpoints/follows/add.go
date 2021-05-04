package follows

import (
	"fmt"
	"github.com/bredbrains/tthk-wish-list/modules"
	"net/http"
	"strconv"
	"time"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var follow models.Follow
	var err error
	err = c.BindJSON(&follow)
	fmt.Println(follow.UserFrom, follow.UserTo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
	}
	follow.CreationTime = time.Now().Format("2006-01-02 15:04:05")
	err, follow = database.AddFollow(follow)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
	}
	message := gin.H{"follow": follow}
	c.JSON(http.StatusOK, message)
}

func ToggleFollowing(c *gin.Context) {
	var follow models.Follow
	var userCalled models.User
	var err error
	var idToConvert int64
	token := c.GetHeader("Token")
	err, userCalled = database.UserData(token)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"success": false, "error": "Invalid token"})
		return
	}
	database.GetFollowsFromUser(userCalled)
	calledId := uint64(userCalled.ID)
	idToConvert, err = strconv.ParseInt(c.Param("id"), 10, 64)
	idForCheck, err := strconv.Atoi(c.Param("id"))
	id := uint64(idToConvert)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "User with this ID doesn't exist."})
		return
	}
	following, isSameUser := modules.CheckIsFollowed(userCalled, idForCheck)
	if isSameUser {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "You cannot start follow yourself."})
		return
	} else {
		follow = models.Follow{UserFrom: calledId, UserTo: id, CreationTime: time.Now().Format("2006-01-02 15:04:05")}
		if following {
			err = database.DeleteFollow(follow)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"success": true})
			return
		} else {
			err, follow = database.AddFollow(follow)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"success": true, "follow": follow})
			return
		}
	}
}
