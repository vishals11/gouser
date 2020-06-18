package controller

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, err error) {
	resp := struct {
		Error string
	}{
		Error: err.Error(),
	}

	json.NewEncoder(w).Encode(resp)
}
