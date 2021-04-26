package database

import (
	"context"
	"database/sql"
	"github.com/bredbrains/tthk-wish-list/models"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func RegisterUser(user models.User) error {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO users(username, first_name, last_name, hash_password, access_token) "+
		"VALUES (?, ?, ?, ?, ?)",
		user.Username, user.FirstName, user.LastName, hash, user.AccessToken)
	return err
}
