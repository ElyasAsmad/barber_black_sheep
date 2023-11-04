package owner_services

import (
	"barber_black_sheep/data"
	"barber_black_sheep/model"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func createService(w http.ResponseWriter, r *http.Request) {
	var service model.Service

	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db, err := sql.Open("sqlite3", data.DB_CONN_STRING)
	if err != nil {
		log.Default().Println(err)
	}
	defer db.Close()
	prepare, err := db.Prepare("INSERT INTO owner_services (service_name, description, duration, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return
	}
	prepare.Exec(service.ServiceName, service.Description, service.Duration, service.Price)
	w.WriteHeader(http.StatusCreated)
	log.Print("Service created successfully")
	//sql query to create a service
	return
}
func listServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//sql query to list all owner_services
	db, err := sql.Open("sqlite3", data.DB_CONN_STRING)
	if err != nil {
		log.Default().Println(err)
	}
	defer db.Close()
	//new slice
	var services []model.Service
	//query
	// owner_services from where owner_id = owner_id left join owner_services on owner_services.service_id = service_owner.service_id
	// select * from owner_services where owner_id = owner_id
	// left join

	rows, err := db.Query("SELECT * FROM owner_services")
	for rows.Next() {
		var service model.Service
		err = rows.Scan(&service.ServiceID, &service.ServiceName, &service.Description, &service.Duration, &service.Price)
		if err != nil {
			log.Default().Println(err)
		}
		if err != nil {
			log.Default().Println(err)
		}
		services = append(services, service)
	}

	if err != nil {
		log.Default().Println(err)
	}

	err = json.NewEncoder(w).Encode(services)
	if err != nil {
		log.Default().Println(err)
	}
	return
}
func getService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//sql query to get a service
	db, err := sql.Open("sqlite3", data.DB_CONN_STRING)
	if err != nil {
		log.Default().Println(err)
	}
	defer db.Close()
	var service model.Service
	serviceID := chi.URLParam(r, "service_id")
	err = db.QueryRow("SELECT * FROM owner_services WHERE service_id = ?", serviceID).Scan(&service.ServiceID, &service.ServiceName, &service.Description, &service.Duration, &service.Price)
	if err != nil {
		log.Default().Println(err)
	}
	err = json.NewEncoder(w).Encode(service)
	if err != nil {
		log.Default().Println(err)
	}
	return
}

func MakeHTTPHandler() http.Handler {
	r := chi.NewRouter()
	r.Post("/", createService)
	r.Get("/", listServices)
	r.Get("/{service_id}", getService)
	return r
}
