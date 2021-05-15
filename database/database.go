package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
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
	db, err = sql.Open("mysql", os.Getenv("MYSQL_CONNECTION_STRING"))
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
		return err
	}
	return err
}

func UserData(accessToken string) (error, models.User) {
	var user models.User
	rows, err := db.Query("SELECT id, email, first_name, last_name FROM users WHERE access_token = ?", accessToken)
	if err != nil {
		return err, user
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName)
	}
	if err != nil {
		return err, user
	}
	return err, user
}

func UserDataById(id int) (error, models.User) {
	var user models.User
	err := db.QueryRow("SELECT id, first_name, last_name FROM users WHERE id = ?", id).Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return err, user
	}
	return err, user
}

func AddWish(wish models.Wish) (error, models.Wish) {
	_, err := db.Exec("INSERT INTO wishes(name, description, user) VALUES(?, ?, ?)", wish.Name, wish.Description, wish.User.ID)
	if err != nil {
		return err, wish
	}
	return err, wish
}

func DeleteWish(wish models.Wish) error {
	_, err := db.Exec("DELETE FROM wishes WHERE id = ?", wish.ID)
	if err != nil {
		return err
	}
	return err
}

func UpdateWish(wish models.Wish) error {
	_, err := db.Exec("UPDATE wishes SET name = ?, description = ? WHERE id = ?", wish.Name, wish.Description, wish.ID)
	if err != nil {
		return err
	}
	return err
}

func HideWish(wish models.Wish) (error, models.Wish) {
	_, err := db.Exec("UPDATE wishes SET hidden = !hidden WHERE id = ?", wish.ID)
	if err != nil {
		return err, wish
	}
	return err, wish
}

func GetFollowsFromUser(user models.User) []models.Follow {
	rows, err := db.Query("SELECT user_to FROM follows WHERE user_from = ?", user.ID)
	var follow models.Follow
	var follows []models.Follow
	if err != nil {
		return follows
	}
	var tick = 0
	for rows.Next() {
		err = rows.Scan(&follow.UserTo)
		if err != nil {
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
		return err, wishes
	}
	var tick = 0
	for rows.Next() {
		err = rows.Scan(&wish.ID)
		if err != nil {
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
		return err, wishes
	}
	var tick = 0
	for rows.Next() {
		err = rows.Scan(&wish.ID, &wish.Name, &wish.Description, &user.ID, &wish.Hidden, &wish.CreationTime)
		if err != nil {
			return err, wishes
		}
		err, wish.User = UserDataById(user.ID)
		wishes = append(wishes, wish)
		tick++
	}
	return err, wishes
}

func AddFollow(follow models.Follow) (error, models.Follow) {
	fmt.Println(follow.UserFrom, follow.UserTo)
	_, err := db.Exec("INSERT INTO follows(user_from, user_to, creation_time) VALUES(?, ?, ?)", follow.UserFrom, follow.UserTo, follow.CreationTime)
	if err != nil {
		return err, follow
	}
	return err, follow
}

func DeleteFollow(follow models.Follow) error {
	_, err := db.Exec("DELETE FROM follows WHERE user_from = ? AND user_to = ? ", follow.UserFrom, follow.UserTo)
	if err != nil {
		return err
	}
	return err
}

func GetWish(id int) (error, models.Wish) {
	var wish models.Wish
	var userID int
	err := db.QueryRow("SELECT * FROM wishes WHERE id = ?", id).Scan(&wish.ID, &wish.Name, &wish.Description, &userID, &wish.Hidden, &wish.CreationTime)
	err, wish.User = UserDataById(userID)
	if err != nil {
		return err, wish
	}
	return err, wish
}

func EditUser(user models.User) (error, models.User) {
	_, err := db.Exec("UPDATE users SET first_name = ?, last_name = ? WHERE id = ?", user.FirstName, user.LastName, user.ID)
	if err != nil {
		return err, user
	}
	return err, user
}

func LikeExist(like models.Like) bool {
	rows, err := db.Query("SELECT id FROM likes WHERE connection = ? AND connection_type = ? AND user = ?", like.Connection, like.ConnectionType, like.User.ID)
	if err != nil {
		return false
	}
	for rows.Next() {
		return true
	}
	return false
}

func GetLike(id int) (error, models.Like) {
	var like models.Like
	var userID int
	err := db.QueryRow("SELECT * FROM likes WHERE id = ?", id).Scan(&like.ID, &like.Connection, &like.ConnectionType, &userID, &like.CreationTime)
	err, like.User = UserDataById(userID)
	if err != nil {
		return err, like
	}
	return err, like
}

func GetLikeId(like models.Like) int {
	var id int
	rows, err := db.Query("SELECT id FROM likes WHERE connection = ? AND connection_type = ? AND user = ?", like.Connection, like.ConnectionType, like.User.ID)
	if err != nil {
		return 0
	}
	for rows.Next() {
		rows.Scan(&id)
		return id
	}
	return 0
}

func AddLike(like models.Like) (error, models.Like) {
	_, err := db.Exec("INSERT INTO likes(connection, connection_type, user, creation_time) VALUES(?, ?, ?, ?)", like.Connection, like.ConnectionType, like.User.ID, like.CreationTime)
	if err != nil {
		return err, like
	}
	return err, like
}

func DeleteLike(like models.Like) error {
	_, err := db.Exec("DELETE FROM likes WHERE id = ?", GetLikeId(like))
	if err != nil {
		return err
	}
	return err
}

func AddComment(comment models.Comment) (error, models.Comment) {
	_, err := db.Exec("INSERT INTO comments(content, connection, connection_type, user, creation_time) VALUES(?, ?, ?, ?, ?)", comment.Content, comment.Connection, comment.ConnectionType, comment.User.ID, comment.CreationTime)
	if err != nil {
		return err, comment
	}
	return err, comment
}
