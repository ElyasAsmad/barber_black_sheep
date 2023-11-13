package main

import (
	admin_owner "barber_black_sheep/api/admin/owner"
	admin_owner_services "barber_black_sheep/api/admin/owner_services"
	admin_user "barber_black_sheep/api/admin/user"
	business_appointment "barber_black_sheep/api/business/owner_appointment"
	business_service "barber_black_sheep/api/business/owner_services"
	public_login "barber_black_sheep/api/public/login"
	user_services "barber_black_sheep/api/user/services"
	"barber_black_sheep/api/user/user_appointment"
	"barber_black_sheep/data"
	"barber_black_sheep/helpers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	log.Default().Println("Server started")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Mount("/api/v1/business/services", business_service.MakeHTTPHandler())
	r.Mount("/api/v1/business/appointments", business_appointment.MakeHttpHandler())
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(helpers.TokenAuth))
		r.Use(data.AdminAuth)
		r.Mount("/api/v1/admin/owners", admin_owner.MakeHTTPHandler())
		r.Mount("/api/v1/admin/users", admin_user.MakeHTTPHandler())
		r.Mount("/api/v1/admin/services", admin_owner_services.MakeHTTPHandler())
		r.Mount("/admin/dashboard", http.StripPrefix("/admin/dashboard", http.FileServer(http.Dir("./web/dashboard"))))
	})
	r.Mount("/api/v1/appointments", user_appointment.MakeHttpHandler())
	r.Mount("/api/v1/services", user_services.MakeHTTPHandler())
	r.Mount("/api/v1/account", public_login.MakeHTTPHandler())
	err := http.ListenAndServe("127.0.0.1:8080", r)
	if err != nil {
		log.Default().Println(err)
	}
}
