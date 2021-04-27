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

func VerifyUser(user models.User) (error, bool) {
	row := db.QueryRow("SELECT hash_password FROM users WHERE username = ?", user.Username)
	var hash_password string
	var err error
	var verified bool
	if err != nil {
		log.Fatal(err)
		return err, false
	}
	err = row.Scan(&hash_password)
	if err != nil {
		log.Fatal(err)
		return err, false
	}
	verified = CheckPasswordHash(user.Password, hash_password)
	return err, verified
}

func RegisterUser(user models.User) error {
	hash, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = db.Exec("INSERT INTO users(username, first_name, last_name, hash_password, access_token, registration_time) "+
		"VALUES (?, ?, ?, ?, ?, ?)",
		user.Username, user.FirstName, user.LastName, hash, user.AccessToken, user.RegistrationTime)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func AddWish(wish models.Wish) error {
	_, err := db.Exec("INSERT INTO wishes(name, description, user) VALUES(?, ?, ?)", wish.Name, wish.Description, wish.User)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func DeleteWish(wish models.Wish) error {
	_, err := db.Exec("DELETE FROM wishes WHERE id = ?", wish.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func EditWish(wish models.Wish) error {
	_, err := db.Exec("UPDATE wishes SET name = ?, description = ? WHERE id = ?", wish.Name, wish.Description, wish.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func HideWish(wish models.Wish) error {
	_, err := db.Exec("UPDATE wishes SET hidden = !hidden WHERE id = ?", wish.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func GetSuggestion(follow models.Follow) (error, []models.Wish) {
	rows, err := db.Query("SELECT * FROM wishes WHERE user = ?", follow.UserTo)
	var wish models.Wish
	var wishes []models.Wish
	if err != nil {
		log.Fatal(err)
		return err, []models.Wish{}
	}
	var tick int = 0
	for rows.Next() {
		err = rows.Scan(&wish.ID, &wish.Name, &wish.Description, &wish.User, &wish.Hidden)
		if err != nil {
			log.Fatal(err)
			return err, []models.Wish{}
		}
		wishes = append(wishes, wish)
		tick++
	}
	return err, wishes
}

func GetWishes(wish models.Wish) (error, []models.Wish) {
	rows, err := db.Query("SELECT * FROM wishes WHERE user = ?", wish.User)
	var wishes []models.Wish
	if err != nil {
		log.Fatal(err)
		return err, []models.Wish{}
	}
	var tick int = 0
	for rows.Next() {
		err = rows.Scan(&wish.ID, &wish.Name, &wish.Description, &wish.User, &wish.Hidden)
		if err != nil {
			log.Fatal(err)
			return err, []models.Wish{}
		}
		wishes = append(wishes, wish)
		tick++
	}
	return err, wishes
}
