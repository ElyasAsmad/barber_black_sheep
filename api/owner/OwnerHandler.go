package owner

import (
	"barber_black_sheep/model"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

// chi http handler routes for owners
func MakeHTTPHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/", listOwners)
	r.Get("/{owner_id}", getOwner)
	r.Post("/", createOwner)
	return r
}

func createOwner(writer http.ResponseWriter, request *http.Request) {
	var owner model.Owner

	err := json.NewDecoder(request.Body).Decode(&owner)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad request"))
		return
	}

	db, err := sql.Open("sqlite3", "./barbar.db")
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("err"))
		return

	}
	defer db.Close()

	prepare, err := db.Prepare("INSERT INTO owner (username, email, passkey, phone) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	prepare.Exec(owner.Username, owner.Email, owner.Password, owner.Phone)
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Owner created successfully"))
	return

}

func getOwner(writer http.ResponseWriter, request *http.Request) {

}

func listOwners(writer http.ResponseWriter, request *http.Request) {

}
