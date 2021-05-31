package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/bredbrains/tthk-wish-list/models"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var (
	db *sql.DB
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
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	if err != nil {
		panic(err)
	}
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
		err = rows.Close()
		return err, user
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName)
	}
	if err != nil {
		err = rows.Close()
		return err, user
	}
	err = rows.Close()
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

func GetUsers() (error, []models.User) {
	var users []models.User
	var user models.User
	rows, err := db.Query("SELECT id, first_name, last_name, email FROM users")
	if err != nil {
		return err, users
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return err, users
		}
		users = append(users, user)
	}
	return err, users
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
	var liked bool
	if err != nil {
		return err, wishes
	}
	var tick = 0
	var count int
	for rows.Next() {
		err = rows.Scan(&wish.ID, &wish.Name, &wish.Description, &user.ID, &wish.Hidden, &wish.CreationTime)
		like := models.Like{
			Connection:     wish.ID,
			ConnectionType: "wishes",
			User:           user,
		}
		err, count = GetLikesCount(like)
		liked = LikeExist(like)
		if err != nil {
			return err, wishes
		}
		wish.Likes = count
		wish.Liked = liked
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
		err = rows.Close()
		return false
	}
	for rows.Next() {
		err = rows.Close()
		return true
	}
	err = rows.Close()
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

func GetLikeByType(id int, recType string) (error, []models.Like) {
	var likes []models.Like
	var like models.Like
	var userID int
	rows, err := db.Query("SELECT * FROM likes WHERE connection = ? AND connection_type = ?", id, recType)
	for rows.Next() {
		rows.Scan(&like.ID, &like.Connection, &like.ConnectionType, &userID, &like.CreationTime)
		err, like.User = UserDataById(userID)
		if err != nil {
			err = rows.Close()
			return err, likes
		}
		likes = append(likes, like)
	}
	err = rows.Close()
	return err, likes
}

func GetLikeId(like models.Like) int {
	var id int
	rows, err := db.Query("SELECT id FROM likes WHERE connection = ? AND connection_type = ? AND user = ?", like.Connection, like.ConnectionType, like.User.ID)
	if err != nil {
		err = rows.Close()
		return 0
	}
	for rows.Next() {
		rows.Scan(&id)
		err = rows.Close()
		return id
	}
	err = rows.Close()
	return 0
}

func GetLikesCount(like models.Like) (error, int) {
	var count int
	rows, err := db.Query("SELECT COUNT(id) FROM likes WHERE connection = ? and connection_type = ?", like.Connection, like.ConnectionType)
	if err != nil {
		err = rows.Close()
		return err, 0
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	err = rows.Close()
	return err, count
}

func UniteLike(like models.Like, liked bool) error {
	var likes int
	rows, err := db.Query("SELECT COUNT(id) FROM likes WHERE connection = ? AND connection_type = ?", like.Connection, like.ConnectionType)
	if err != nil {
		err = rows.Close()
		return err
	}
	for rows.Next() {
		rows.Scan(&likes)
	}
	_, err = db.Exec("UPDATE "+like.ConnectionType+" SET liked = ?, likes = ? WHERE id = ?", liked, likes, like.Connection)
	if err != nil {
		err = rows.Close()
		return err
	}
	err = rows.Close()
	return err
}

func AddLike(like models.Like) (error, models.Like) {
	rows, err := db.Query("INSERT INTO likes(connection, connection_type, user, creation_time) VALUES(?, ?, ?, ?)", like.Connection, like.ConnectionType, like.User.ID, like.CreationTime)
	if err != nil {
		err = rows.Close()
		return err, like
	}
	err = rows.Close()
	return err, like
}

func DeleteLike(like models.Like) error {
	rows, err := db.Query("DELETE FROM likes WHERE id = ?", GetLikeId(like))
	if err != nil {
		err = rows.Close()
		return err
	}
	err = rows.Close()
	return err
}

func GetComment(id int) (error, models.Comment) {
	var comment models.Comment
	var userID int
	err := db.QueryRow("SELECT * FROM comments WHERE id = ?", id).Scan(&comment.ID, &comment.Content, &comment.Connection, &comment.ConnectionType, &userID, &comment.CreationTime)
	err, comment.User = UserDataById(userID)
	if err != nil {
		return err, comment
	}
	return err, comment
}

func AddComment(comment models.Comment) (error, models.Comment) {
	_, err := db.Exec("INSERT INTO comments(content, connection, connection_type, user, creation_time) VALUES(?, ?, ?, ?, ?)", comment.Content, comment.Connection, comment.ConnectionType, comment.User.ID, comment.CreationTime)
	if err != nil {
		return err, comment
	}
	return err, comment
}

func UpdateComment(comment models.Comment) (error, models.Comment) {
	_, err := db.Exec("UPDATE comments SET content = ?, connection = ?, connection_type = ?, user = ?, creation_time = ?", comment.Content, comment.Connection, comment.ConnectionType, comment.User.ID, comment.CreationTime)
	if err != nil {
		return err, comment
	}
	return err, comment
}

func DeleteComment(comment models.Comment) error {
	_, err := db.Exec("DELETE FROM comments WHERE id = ?", comment.ID)
	if err != nil {
		return err
	}
	return err
}
