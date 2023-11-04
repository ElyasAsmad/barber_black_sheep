package main

import (
	"barber_black_sheep/api/login"
	"barber_black_sheep/api/owner"
	"barber_black_sheep/api/owner_appointment"
	"barber_black_sheep/api/owner_services"
	"barber_black_sheep/api/services"
	"barber_black_sheep/api/user"
	"barber_black_sheep/api/user_appointment"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {

	log.Default().Println("Server started")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api/v1/business/owner_services", owner_services.MakeHTTPHandler())
	r.Mount("/api/v1/business/appointments", owner_appointment.MakeHttpHandler())
	r.Mount("/api/v1/admin/owners", owner.MakeHTTPHandler())
	r.Mount("/api/v1/admin/users", user.MakeHTTPHandler())
	r.Mount("/api/v1/appointments", user_appointment.MakeHttpHandler())
	r.Mount("/api/v1/services", services.MakeHTTPHandler())
	r.Mount("/api/v1/account", login.MakeHTTPHandler())
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Default().Println(err)
	}
}
