package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email, omitempty"`
}

func helloWriter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func personWriter(w http.ResponseWriter, r *http.Request) {
	user0 := Person{
		Name:  "Jack",
		Age:   54,
		Email: "jack@gmain.com",
	}
	resp, err := json.Marshal(user0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

func main() {

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
	http.HandleFunc("/", helloWriter)
	http.HandleFunc("/person", personWriter)
}
