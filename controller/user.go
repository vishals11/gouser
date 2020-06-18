package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vishals11/gouser/model"
	"gopkg.in/go-playground/validator.v9"
)

// using a single instance of Validate as it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		err = fmt.Errorf("Error Decoding Input:%s", err)
		WriteError(w, err)
		return
	}

	err = validate.Struct(user)
	if err != nil {
		err = fmt.Errorf("Validation failed:%s", err)
		WriteError(w, err)
		return
	}

	err = model.CreateUser(&user)
	if err != nil {
		err = fmt.Errorf("Error while creating new user: %s", err)
		WriteError(w, err)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var login model.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		err = fmt.Errorf("Error Decoding Input:%s", err)
		WriteError(w, err)
		return
	}

	err = validate.Struct(login)
	if err != nil {
		err = fmt.Errorf("validation failed:%v", err)
		WriteError(w, err)
		return
	}

	user, err := model.LoginUser(login)
	if err != nil {
		err = fmt.Errorf("Error during Login: %s", err)
		WriteError(w, err)
		return
	}
	json.NewEncoder(w).Encode(user)
}
