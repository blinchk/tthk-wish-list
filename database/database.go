package database

import (
	"context"
	"database/sql"
	"fmt"
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

func VerifyUser(user models.User) bool {
	row := db.QueryRow("SELECT hash_password FROM users WHERE username = ?", user.Username)
	var hash_password string
	switch err := row.Scan(&hash_password); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return false
	case nil:
		fmt.Println(hash_password)
		if CheckPasswordHash(user.Password, hash_password) {
			return true
		}
	default:
		panic(err)
	}
	return false
}

func RegisterUser(user models.User) error {
	hash, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = db.Exec("INSERT INTO users(username, first_name, last_name, hash_password, access_token) "+
		"VALUES (?, ?, ?, ?, ?)",
		user.Username, user.FirstName, user.LastName, hash, user.AccessToken)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func AddWish(wish models.Wish) error {
	_, err := db.Exec("INSERT INTO wishes(name, description) VALUES(?, ?)", wish.Name, wish.Description)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func DeleteWish(wish models.Wish) error {
	_, err := db.Exec("DELETE FROM wishes WHERE name = ? AND description = ?", wish.Name, wish.Description)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
