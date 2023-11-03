package model

type Availability struct {
	AvailabilityID string `json:"availability_id"` // primary key
	OwnerID        string `json:"owner_id"`        // foreign key
	ServiceID      string `json:"service_id"`      // foreign key
	DayOfWeek      string `json:"day_of_week"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
}

// sqlite query to create table
