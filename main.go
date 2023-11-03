package main

import (
	"barber_black_sheep/api/owner"
	"barber_black_sheep/api/services"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Mount("/api/v1/business/services", services.MakeHTTPHandler())
	r.Mount("api/v1/admin/owners", owner.MakeHTTPHandler())
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Default().Println(err)
	}
}
