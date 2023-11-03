package model

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

//sqlite query to create table
// CREATE TABLE user (
// 	user_id INTEGER PRIMARY KEY AUTOINCREMENT,
// 	username TEXT,
// 	email TEXT,
// 	password TEXT,
// 	phone TEXT
// );
