package model

import (
	"barber_black_sheep/data"
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email" `
	Password string `json:"password"`
	Phone    string `json:"phone"`
	// Role     string `json:"role"` // user or owner
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CreateUser(user *User) error {

	log.Default().Println(user.Username, user.Email, user.Password, user.Phone)
	existUser := GetUserByUsername(user.Username)
	if existUser != nil {
		log.Println(existUser.Username)
		return errors.New("user already exists")
	}
	db, err := sql.Open("sqlite3", data.DB_CONN_STRING)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer db.Close()

	prepare, err := db.Prepare("INSERT INTO users (full_name, username, email, password, phone) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	//hash user password
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	prepare.Exec(user.FullName, user.Username, user.Email, user.Password, user.Phone)
	return nil
}

// check if user exists
// if user exists, return user
// if user does not exist, return nil
func GetUserByUsername(username string) *User {
	db, err := sql.Open("sqlite3", data.DB_CONN_STRING)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer db.Close()
	var user User
	log.Println(username, "Creating query")

	row := db.QueryRow("SELECT * FROM users WHERE username = ?", username)
	err = row.Scan(&user.UserID, &user.FullName, &user.Username, &user.Email, &user.Password, &user.Phone)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	err = row.Err()
	if err != nil {
		log.Println(err.Error())

		return nil
	}
	return &user
}

//sqlite query to create table
// CREATE TABLE user (
// 	user_id INTEGER PRIMARY KEY AUTOINCREMENT,
// 	full_name TEXT,
// 	username TEXT,
// 	email TEXT,
// 	password TEXT,
// 	phone TEXT
// );
