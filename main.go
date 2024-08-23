package main

import (
	"HandlersTask/user_handlers"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func helloWriter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Invalid Request"))
}

func main() {

	http.HandleFunc("/", helloWriter)
	http.HandleFunc("/api/get_users", user_handlers.GetUsers)
	http.HandleFunc("/api/get_one_user", user_handlers.GetOneUser)
	http.HandleFunc("/api/delete_user", user_handlers.DeleteUser)
	http.HandleFunc("/api/create_user", user_handlers.CreateUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server", err)
		return
	}
}
