package wishes

import (
	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
)

func AssignUserToWish(wish models.Wish, token string) (error, models.Wish) {
	var err error
	err, wish.User = database.UserData(token)
	if err != nil {
		return err, wish
	}
	return err, wish
}

func CheckWishPermissions(wish models.Wish, token string) (error, bool) {
	var err error
	var user models.User
	var currentWish models.Wish
	if token == "" {
		return err, false
	}
	err, user = database.UserData(token)
	if err != nil {
		return err, false
	}
	err, currentWish = database.GetWish(wish.ID)
	if currentWish.User.ID == user.ID {
		return err, true
	}
	return err, false
}
