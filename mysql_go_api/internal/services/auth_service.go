package services

import (
	"encoding/json"
	"net/http"

	"github.com/Utsavch189/api_mysql/internal/controller"
	"github.com/Utsavch189/api_mysql/internal/models/requests"
	"github.com/Utsavch189/api_mysql/internal/models/responses"
	"github.com/Utsavch189/api_mysql/internal/utils"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser requests.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.ErrorResponse(err, "user creation failed!"))
		return
	}
	user, uerr := requests.NewUser(newUser.UserName, newUser.Password)
	if uerr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.ErrorResponse(uerr, "user creation failed!"))
		return
	}
	if err := utils.ValidateStruct(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.ErrorResponse(err, "user creation failed!", utils.FormatValidationError(err)))
		return
	}
	createdUser, cerr := controller.Register(user)
	if cerr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.ErrorResponse(cerr, "user creation failed!"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses.CreateUserResponse(createdUser, "user creation successful!"))
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var login requests.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.ErrorResponse(err, "login failed!"))
		return
	}
	if err := utils.ValidateStruct(login); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.ErrorResponse(err, "login failed!", utils.FormatValidationError(err)))
		return
	}
	user, uerr := controller.GetAUserByUsername(login.UserName)
	if uerr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.ErrorResponse(uerr, "login failed!"))
		return
	}

	verr := utils.VerifyHash(user.Password, login.Password)

	if verr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.ErrorResponse(verr, "wrong password!"))
		return
	}

	token, jerr := utils.GenerateJWT(user.Id, user.UserName)

	if jerr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.ErrorResponse(jerr, "jwt creation failed!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses.LoginResponse(token, "login successful!"))
}
