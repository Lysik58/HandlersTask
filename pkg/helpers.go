package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, error, details string, status int) {
	resp, err := json.Marshal(struct {
		Title   string `json:"title"`
		Error   string `json:"error"`
		Details string `json:"details"`
		Status  int    `json:"status"`
	}{
		Title:   "Error",
		Error:   error,
		Details: details,
		Status:  status,
	})
	if err != nil {
		log.Fatal("Error func error response", err)
	}
	_, err = w.Write(resp)
	if err != nil {
		log.Fatal("Error func error response", err)
	}

}
