package login

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

//chi router login

func MakeHTTPHandler() http.Handler {
	r := chi.NewRouter()
	r.Post("/", Login)
	r.Post("/logout", Logout)
	return r
}
