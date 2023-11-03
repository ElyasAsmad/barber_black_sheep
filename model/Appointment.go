package model

type Appointment struct {
	AppointmentID int    `json:"appointment_id"` // primary key
	UserID        int    `json:"user_id"`        // foreign key
	ServiceID     int    `json:"service_id"`     // foreign key
	Date          string `json:"date"`
	Time          string `json:"time"`
	Status        string `json:"status"`
}

//sqlite query to create table
