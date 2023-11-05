package model

import (
	"barber_black_sheep/data"
	"database/sql"
	"log"
)

type Role struct {
	UserId   int    `json:"user_id"` // primary key from users table
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}

func GetRoleByUser(user_id int) (Role, error) {
	var role Role
	db, err := sql.Open("sqlite3", data.DB_CONN_STRING)
	if err != nil {
		log.Println(err.Error())
	}
	res := db.QueryRow("SELECT * FROM role WHERE user_id = ?", user_id)
	err = res.Scan(&role.UserId, &role.RoleId, &role.RoleName)
	if err != nil {
		log.Println(err.Error())
	}
	return role, err
}
