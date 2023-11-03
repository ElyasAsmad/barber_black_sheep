package enum

// AppointmentStatusEnum is a string that represents the status of an appointment
type AppointmentStatusEnum string

const (
	// Pending is an appointment that has not been confirmed by the business owner
	Pending AppointmentStatusEnum = "pending"
	// Confirmed is an appointment that has been confirmed by the business owner
	Confirmed AppointmentStatusEnum = "confirmed"
	// Completed is an appointment that has been completed by the business owner
	Completed AppointmentStatusEnum = "completed"
	// Cancelled is an appointment that has been cancelled by the business owner
	Cancelled AppointmentStatusEnum = "cancelled"
)
