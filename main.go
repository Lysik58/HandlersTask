package main

import (
	"HandlersTask/UserHandlers"
	"net/http"
)

func helloWriter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {

	http.HandleFunc("/", helloWriter)
	http.HandleFunc("/person", UserHandlers.PersonWriter)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}
