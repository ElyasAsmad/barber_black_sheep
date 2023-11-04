package model

type Service struct {
	ServiceID   int     `json:"service_id"`
	ServiceName string  `json:"service_name"`
	Description string  `json:"description"`
	Duration    string  `json:"duration"`
	Price       float64 `json:"price"`
}
