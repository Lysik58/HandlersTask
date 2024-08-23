package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
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

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", DbInfo)
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, err
}

func AddHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
