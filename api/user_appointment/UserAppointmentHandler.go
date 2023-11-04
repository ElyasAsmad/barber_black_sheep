package user_appointment

import (
	"barber_black_sheep/model"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// chi http handler routes for user_appointment
// user can create, update, delete, and list appointments

func MakeHttpHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/", GetAllAppointments)
	r.Get("/{id}", GetAppointment)
	r.Post("/", CreateAppointment)
	r.Put("/{id}", UpdateAppointment)
	r.Delete("/{id}", DeleteAppointment)
	return r
}

func DeleteAppointment(writer http.ResponseWriter, request *http.Request) {

}

func UpdateAppointment(writer http.ResponseWriter, request *http.Request) {

}

func CreateAppointment(writer http.ResponseWriter, request *http.Request) {
	db, err := sql.Open("sqlite3", "./barbar.db")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	defer db.Close()
	var appointment model.Appointment
	err = json.NewDecoder(request.Body).Decode(&appointment)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}
	prepare, err := db.Prepare("INSERT INTO appointments (service_id, user_id, date, time) VALUES (?, ?, ?, ?)")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	prepare.Exec(appointment.ServiceID, appointment.UserID, appointment.Date, appointment.Time)
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Appointment created successfully"))
	return

}

func GetAppointment(writer http.ResponseWriter, request *http.Request) {

}

func GetAllAppointments(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// get all appointments
	db, err := sql.Open("sqlite3", "./barbar.db")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	defer db.Close()
	var appointments []model.Appointment
	rows, err := db.Query("SELECT * FROM appointments")
	for rows.Next() {
		var appointment model.Appointment
		err = rows.Scan(&appointment.AppointmentID, &appointment.ServiceID, &appointment.UserID, &appointment.Date, &appointment.Time)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}
		appointments = append(appointments, appointment)
	}
	err = json.NewEncoder(writer).Encode(appointments)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	return
}
