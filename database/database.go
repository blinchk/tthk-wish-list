package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/bredbrains/tthk-wish-list/models"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var (
	ctx context.Context
	db  *sql.DB
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Connect() {
	var err error
	db, err = sql.Open("mysql", "wish_list:LVA7ECV3ucv3MGDp@tcp(mysql-db.bredbrains.tech)/wish_list")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

func VerifyUser(user models.User) (error, bool, string) {
	var hashPassword string
	var accessToken string
	var verified bool
	rows, err := db.Query("SELECT hash_password, access_token FROM users WHERE email = ?", user.Email)
	for rows.Next() {
		rows.Scan(&hashPassword, &accessToken)
	}
	if err != nil {
		log.Fatal(err)
		return err, false, ""
	}
	verified = CheckPasswordHash(user.Password, hashPassword)
	return err, verified, accessToken
}

func RegisterUser(user models.User) error {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO users(email, first_name, last_name, hash_password, access_token, registration_time) "+
		"VALUES (?, ?, ?, ?, ?, ?)",
		user.Email, user.FirstName, user.LastName, hash, user.AccessToken, user.RegistrationTime)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func UserData(accessToken string) (error, models.User) {
	var user models.User
	rows, err := db.Query("SELECT id, email, first_name, last_name FROM users WHERE access_token = ?", accessToken)
	if err != nil {
		log.Fatal(err)
		return err, user
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName)
	}
	if err != nil {
		log.Fatal(err)
		return err, user
	}
	return err, user
}

func UserDataById(id int) (error, models.User) {
	var user models.User
	rows, err := db.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
		return err, user
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName)
	}
	if err != nil {
		log.Fatal(err)
		return err, user
	}
	return err, user
}

func AddWish(wish models.Wish) (error, models.Wish) {
	_, err := db.Exec("INSERT INTO wishes(name, description, user) VALUES(?, ?, ?)", wish.Name, wish.Description, wish.User.ID)
	if err != nil {
		log.Fatal(err)
		return err, wish
	}
	return err, wish
}

func DeleteWish(wish models.Wish) (error, models.Wish) {
	_, err := db.Exec("DELETE FROM wishes WHERE id = ?", wish.ID)
	if err != nil {
		log.Fatal(err)
		return err, wish
	}
	return err, wish
}

func UpdateWish(wish models.Wish) (error, models.Wish) {
	_, err := db.Exec("UPDATE wishes SET name = ?, description = ? WHERE id = ?", wish.Name, wish.Description, wish.ID)
	if err != nil {
		log.Fatal(err)
		return err, wish
	}
	return err, wish
}

func HideWish(wish models.Wish) (error, models.Wish) {
	_, err := db.Exec("UPDATE wishes SET hidden = !hidden WHERE id = ?", wish.ID)
	if err != nil {
		log.Fatal(err)
		return err, wish
	}
	return err, wish
}

func GetFollowsFromUser(user models.User) []models.Follow {
	rows, err := db.Query("SELECT user_to FROM follows WHERE user_from = ?", user.ID)
	var follow models.Follow
	var follows []models.Follow
	if err != nil {
		log.Fatal(err)
		return follows
	}
	var tick int = 0
	for rows.Next() {
		err = rows.Scan(&follow.UserTo)
		if err != nil {
			log.Fatal(err)
			return follows
		}
		follows = append(follows, follow)
		tick++
	}
	return follows
}

func GetSuggestion(follow models.Follow) (error, []models.Wish) {
	rows, err := db.Query("SELECT id FROM wishes WHERE user = ?", follow.UserTo)
	var wish models.Wish
	var wishes []models.Wish
	if err != nil {
		log.Fatal(err)
		return err, wishes
	}
	var tick int = 0
	for rows.Next() {
		err = rows.Scan(&wish.ID)
		if err != nil {
			log.Fatal(err)
			return err, wishes
		}
		wishes = append(wishes, wish)
		tick++
	}
	return err, wishes
}

func GetWishes(user models.User) (error, []models.Wish) {
	rows, err := db.Query("SELECT * FROM wishes WHERE user = ?", user.ID)
	var wishes []models.Wish
	var wish models.Wish
	if err != nil {
		log.Fatal(err)
		return err, wishes
	}
	var tick int = 0
	for rows.Next() {
		err = rows.Scan(&wish.ID, &wish.Name, &wish.Description, &user.ID, &wish.Hidden, &wish.CreationTime)
		if err != nil {
			log.Fatal(err)
			return err, wishes
		}
		err, wish.User = UserDataById(user.ID)
		wishes = append(wishes, wish)
		tick++
	}
	return err, wishes
}

func AddFollow(follow models.Follow) (error, models.Follow) {
	_, err := db.Exec("INSERT INTO follows(user_from, user_to, creation_time) VALUES(?, ?, ?)", follow.UserFrom, follow.UserTo, follow.CreationTime)
	if err != nil {
		log.Fatal(err)
		return err, follow
	}
	return err, follow
}

func DeleteFollow(follow models.Follow) (error, models.Follow) {
	_, err := db.Exec("DELETE FROM follows WHERE id = ?", follow.ID)
	if err != nil {
		log.Fatal(err)
		return err, follow
	}
	return err, follow
}
