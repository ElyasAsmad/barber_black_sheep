package admin_owner_appointment

import (
	"barber_black_sheep/data"
	"barber_black_sheep/model"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func MakeHttpHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/", GetAllAppointments)
	r.Get("/{id}", GetAppointment)
	r.Put("/{id}", UpdateAppointment)
	return r
}

func UpdateAppointment(writer http.ResponseWriter, request *http.Request) {

}

func GetAppointment(writer http.ResponseWriter, request *http.Request) {
	// get appointment by id
	res := chi.URLParam(request, "id")
	db, err := sql.Open("sqlite3", data.DB_CONN_STRING)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	defer db.Close()
	var appointment model.Appointment
	err = db.QueryRow("SELECT * FROM appointment WHERE appointment_id = ?", res).Scan(&appointment.AppointmentID, &appointment.ServiceID, &appointment.UserID, &appointment.Date, &appointment.Time)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	err = json.NewEncoder(writer).Encode(appointment)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	return
}

func GetAllAppointments(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// get all appointments
	db, err := sql.Open("sqlite3", data.DB_CONN_STRING)
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
