package model

type Owner struct {
	OwnerID  int    `json:"owner_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
} // Owner has many appointments

// Owner has many owner_services

//turn struct into json literals
//json literals are key value pairs
// { "key": "value
