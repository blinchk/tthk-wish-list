package router

import (
	"github.com/bredbrains/tthk-wish-list/endpoints/feedback"
	"github.com/bredbrains/tthk-wish-list/endpoints/feedback/comments"
	"github.com/bredbrains/tthk-wish-list/endpoints/wishes"
	"github.com/gin-gonic/gin"
)

func AddWishEndpoints(router *gin.Engine) {
	wishAPI := router.Group("/wish")
	wishAPI.GET("/suggestion/:id", isAuthorized(wishes.GetSuggestion))
	wishAPI.PUT("/", isAuthorized(wishes.Add))
	wishAPI.DELETE("/:id", isAuthorized(wishes.Delete))
	wishAPI.PATCH("/", isAuthorized(wishes.Update))
	wishAPI.PATCH("/:id/hide", isAuthorized(wishes.Hide))
	wishAPI.POST("/like", isAuthorized(feedback.ToggleLike))
	wishAPI.GET("/:id/likes", isAuthorized(feedback.GetLikesByWish))
	wishAPI.GET("/:id/:type/likes", isAuthorized(feedback.GetLikesCount))
	wishAPI.POST("/comment", isAuthorized(comments.Add))
	wishAPI.PATCH("/:id/comment", isAuthorized(comments.Update))
	wishAPI.DELETE("/:id/comment", isAuthorized(comments.Delete))
	wishAPI.GET("/:id/comment", isAuthorized(comments.GetComment))
	wishAPI.POST("/gift/:wish", isAuthorized(feedback.ToggleGift))
	wishAPI.GET("/:id/gift", isAuthorized(feedback.GetGiftByWish))
	wishAPI.GET("/gift/:id", isAuthorized(feedback.GetGiftsByUser))
	wishAPI.PATCH("/gift/", isAuthorized(feedback.EditGift))
	wishAPI.GET("/gifts", isAuthorized(feedback.GetGifts))
	wishAPI.POST("/gift/:wish/book", isAuthorized(feedback.ToogleBookingByWish))
}
