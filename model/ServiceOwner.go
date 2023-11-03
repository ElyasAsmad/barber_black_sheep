package model

type ServiceOwner struct {
	ServiceOwnerID int `json:"service_owner_id"` // primary key
	OwnerID        int `json:"owner_id"`         // foreign key
	ServiceID      int `json:"service_id"`       // foreign key
}

//sqlite query to create table
