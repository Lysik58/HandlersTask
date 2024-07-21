package main

import (
	"HandlersTask/UserHandlers"
	"fmt"
	"log"
	"net/http"
)

func helloWriter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Invalid Request"))
}

func main() {

	http.HandleFunc("/", helloWriter)
	http.HandleFunc("/api/get_users", UserHandlers.GetUsers)
	http.HandleFunc("/api/get_one_user", UserHandlers.GetOneUser)
	http.HandleFunc("/api/delete_user", UserHandlers.DeleteUser)
	http.HandleFunc("/api/create_user", UserHandlers.CreateUser)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running...")
	if err != nil {
		log.Fatal("Error starting server", err)
		return
	}

}
