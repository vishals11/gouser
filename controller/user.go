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
		err = fmt.Errorf("Validation failed:%s\n", err)
		WriteError(w, err)
		return
	}

	resp, err := model.CreateUser(&user)
	if err != nil {
		err = fmt.Errorf("Error while creating new user: %s", err)
		WriteError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
