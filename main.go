package main

import (
	"barber_black_sheep/api/owner"
	"barber_black_sheep/api/services"
	"barber_black_sheep/api/user"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	log.Default().Println("Server started")
	r := chi.NewRouter()

	r.Mount("/api/v1/business/services", services.MakeHTTPHandler())
	r.Mount("/api/v1/admin/owners", owner.MakeHTTPHandler())
	r.Mount("/api/v1/admin/users", user.MakeHTTPHandler())
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Default().Println(err)
	}
}
