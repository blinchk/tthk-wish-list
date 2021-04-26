package database

import (
	"context"
	"database/sql"
	"github.com/bredbrains/tthk-wish-list/models"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
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

func RegisterUser(user models.User, db sql.DB) bool {
	hash, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err = db.ExecContext(ctx, "INSERT INTO users(username, hash_password, access_token) VALUES (?, ?, ?)", user.Username, hash)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
