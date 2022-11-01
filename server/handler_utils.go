package server

import (
	"encoding/json"
	"net/http"
)

func httpError(w http.ResponseWriter, status int, message string) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	responseJSON, err := json.Marshal(map[string]string{"error": message})
	if err != nil {
		return err
	}

	_, err = w.Write(responseJSON)

	return err
}

func httpSuccess(w http.ResponseWriter, rp relayProperties) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	responseJSON, err := json.Marshal(rp)
	if err != nil {
		return err
	}

	_, err = w.Write(responseJSON)

	return err
}
