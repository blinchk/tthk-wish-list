package follows

import (
	"fmt"
	"net/http"
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
