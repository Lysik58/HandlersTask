package UserHandlers

import (
	"encoding/json"
	"fmt"
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
		fmt.Println(err)
	}
	_, err = w.Write(resp)
	if err != nil {
		fmt.Println(err)
	}

}
