package modules

import (
	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
)

func CheckIsFollowed(requestingUser models.User, id int) (bool, bool) {
	var following, isSameUser bool
	if requestingUser.ID != id {
		follows := database.GetFollowsFromUser(requestingUser)
		for _, follow := range follows {
			if int(follow.UserTo) == id {
				following = true
				break
			}
		}
	} else {
		following = false
		isSameUser = true
	}
	return following, isSameUser
}
