package model

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email" `
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"` // user or owner
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CreateUser(user *User) error {
	//open database
	// defer db.Close()
	//hash password
	//prepare statement
	//execute statement
	//return nil or error
	log.Default().Println(user.Username, user.Email, user.Password, user.Phone, user.Role)
	existUser := GetUserByUsername(user.Username)
	if existUser != nil {
		log.Println(existUser.Username)
		return errors.New("user already exists")
	}
	db, err := sql.Open("sqlite3", "./barbar.db?_busy_timeout=5000")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer db.Close()

	prepare, err := db.Prepare("INSERT INTO users (username, email, password, phone, role) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	//hash user password
	user.Password, err = HashPassword(user.Password)
	prepare.Exec(user.Username, user.Email, user.Password, user.Phone, user.Role)
	return nil
}

// check if user exists
// if user exists, return user
// if user does not exist, return nil
func GetUserByUsername(username string) *User {
	db, err := sql.Open("sqlite3", "./barbar.db?_busy_timeout=5000")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer db.Close()
	var user User
	log.Println(username, "Creating query")

	row := db.QueryRow("SELECT * FROM users WHERE username = ?", username)
	err = row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Phone, &user.Role)
	if err != nil {
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
// 	username TEXT,
// 	email TEXT,
// 	password TEXT,
// 	phone TEXT
// );
