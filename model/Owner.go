package model

type Owner struct {
	OwnerID  int    `json:"owner_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
