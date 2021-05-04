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
