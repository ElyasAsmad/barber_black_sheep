package user_services

import (
	"barber_black_sheep/data"
	"barber_black_sheep/model"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/", GetAllServices)
	r.Get("/{id}", GetService)
	return r
}

func GetService(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	res := chi.URLParam(request, "id")
	db, err := sql.Open("sqlite3", data.DB_CONN_STRING)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("err"))
		return
	}
	defer db.Close()
	var service model.Service
	rows, err := db.Query("SELECT * FROM services WHERE service_id = ?", res)
	for rows.Next() {
		err = rows.Scan(&service.ServiceID, &service.ServiceName, &service.Description, &service.Duration, &service.Price)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("err"))
			return
		}

	}
	err = rows.Err()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	if service.ServiceID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Service not found"))
		return
	} else {
		json.NewEncoder(writer).Encode(service)
		return
	}
}

func GetAllServices(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("sqlite3", data.DB_CONN_STRING)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("err"))
		return
	}
	defer db.Close()
	var services []model.Service
	rows, err := db.Query("SELECT * FROM services")
	for rows.Next() {
		var service model.Service
		err = rows.Scan(&service.ServiceID, &service.ServiceName, &service.Description, &service.Duration, &service.Price)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("err"))
			return
		}
		services = append(services, service)
	}
	err = rows.Err()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	err = json.NewEncoder(writer).Encode(services)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("err"))
		return
	}
	return
}
