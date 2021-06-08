package router

import (
	"github.com/bredbrains/tthk-wish-list/endpoints/auth"
	"github.com/bredbrains/tthk-wish-list/endpoints/follows"
	"github.com/bredbrains/tthk-wish-list/endpoints/users"
	"github.com/gin-gonic/gin"
)

func AddUserEndpoints(router *gin.Engine) {
	userAPI := router.Group("/user")
	userAPI.GET("/", isAuthorized(auth.User))
	userAPI.PATCH("/", isAuthorized(users.EditUserProfile))
	userAPI.GET("/users", users.GetUsersProfiles)
	userAPI.GET("/:id", users.GetUserProfile)
	userAPI.GET("/:id/wishes", isAuthorized(users.Wishes))
	userAPI.POST("/:id/follow", isAuthorized(follows.ToggleFollowing))
}
