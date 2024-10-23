package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// Menangani error dan menghentikan aplikasi jika error kritis
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Kirim respons dalam format JSON
func SendResponse(w http.ResponseWriter, status int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		SendError(w, http.StatusInternalServerError, "Failed to encode JSON")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// Kirim pesan error dalam format JSON
func SendError(w http.ResponseWriter, status int, message string) {
	SendResponse(w, status, map[string]string{"error": message})
}

// Parsing body permintaan HTTP ke struct
func DecodeBody(r *http.Request, target interface{}) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(target)
	return err
}
