package main

import (
	"HandlersTask/UserHandlers"
	"log"
	"net/http"
)

func helloWriter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {

	http.HandleFunc("/", helloWriter)
	http.HandleFunc("/api/getusers", UserHandlers.GetUsers)
	http.HandleFunc("/api/getoneuser", UserHandlers.GetOneUser)
	http.HandleFunc("/api/deleteuser", UserHandlers.DeleteUser)
	http.HandleFunc("/api/createuser", UserHandlers.CreateUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server", err)
		return
	}

}
