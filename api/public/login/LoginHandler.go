package public_login

import (
	"barber_black_sheep/enum"
	"barber_black_sheep/helpers"
	"barber_black_sheep/model"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

//chi router login

func MakeHTTPHandler() http.Handler {
	r := chi.NewRouter()
	r.Post("/login", Login)
	r.Post("/logout", Logout)
	r.Post("/register", Register)
	return r
}

func Register(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad request"))
		return
	}
	user.Role = strconv.Itoa(int(enum.User))
	err = model.CreateUser(&user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.WriteHeader(http.StatusCreated)
	return
}

func Logout(writer http.ResponseWriter, request *http.Request) {

}

func Login(writer http.ResponseWriter, request *http.Request) {
	var login model.LoginRequest
	err := json.NewDecoder(request.Body).Decode(&login)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad request"))
		return
	}
	userRes := model.GetUserByUsername(login.Username)
	if userRes == nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("User not found"))
		return
	}
	passResult := helpers.ComparePassword(login.Password, userRes.Password)

	if userRes.Username == "" || !passResult {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("User not found"))
		return
	}
	token, err := helpers.GenerateJWT(*userRes)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error generating token"))
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(token))
	//some logic to check if user exists
	//if user exists, check if password matches
	//if password matches, create a jwt token
	//return jwt token
	//else return error
}
