package user

import (
	"barber_black_sheep/data"
	"barber_black_sheep/enum"
	"barber_black_sheep/helpers"
	"barber_black_sheep/model"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"strconv"
)

func MakeHTTPHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(helpers.TokenAuth))
	r.Use(jwtauth.Authenticator)
	r.Get("/", listUsers)
	r.Get("/{user_id}", getUser)
	r.Post("/", createUser)
	return r
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	var user model.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad request"))
		return
	}

	switch user.Role {
	case "admin":
		user.Role = strconv.Itoa(int(enum.Admin))
	case "owner":
		user.Role = strconv.Itoa(int(enum.Owner))
	default:
		user.Role = strconv.Itoa(int(enum.User))
	}
	err = model.CreateUser(&user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.WriteHeader(http.StatusCreated)
	return
}

func listUsers(writer http.ResponseWriter, request *http.Request) {
	db, err := sql.Open("sqlite3", data.DB_CONN_STRING)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("err"))
		return
	}
	defer db.Close()
	var users []model.User
	rows, err := db.Query("SELECT * FROM users")
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Phone, &user.Role)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}
		users = append(users, user)
	}
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	err = json.NewEncoder(writer).Encode(users)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.WriteHeader(http.StatusOK)
	return
}

func getUser(writer http.ResponseWriter, request *http.Request) {
	//chi url param
	userID := chi.URLParam(request, "user_id")
	db, err := sql.Open("sqlite3", data.DB_CONN_STRING)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("err"))
		return
	}
	defer db.Close()
	var user model.User
	err = db.QueryRow("SELECT * FROM users WHERE user_id = ?", userID).Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Phone, &user.Role)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("err"))
		return
	}
	err = json.NewEncoder(writer).Encode(user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("err"))
		return
	}
	writer.WriteHeader(http.StatusOK)
	return

}
