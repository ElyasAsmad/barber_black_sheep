package model

// Service is a struct that represents a service
// offered by a business owner.
//create struct
//
// type Service struct {
// 	ServiceID   int     `json:"service_id"`
// 	ServiceName string  `json:"service_name"`
// 	Description string  `json:"description"`
// 	Duration    string  `json:"duration"`
// 	Price       float64 `json:"price"`
// }

type Service struct {
	ServiceID   int     `json:"service_id"`
	ServiceName string  `json:"service_name"`
	Description string  `json:"description"`
	Duration    string  `json:"duration"`
	Price       float64 `json:"price"`
}

//sql query to create table
