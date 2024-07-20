package UserHandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Person struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

var PersonData = map[int]Person{
	1: {
		Id:    1,
		Name:  "John",
		Age:   32,
		Email: "john@gmail.com",
	},
	2: {
		Id:    2,
		Name:  "Sam",
		Age:   32,
		Email: "SeriousSam@gmail.com",
	},
	3: {
		Id:    3,
		Name:  "Emily",
		Age:   32,
		Email: "emily@gmail.com",
	},
	4: {
		Id:    4,
		Name:  "Doctor",
		Age:   90,
		Email: "WhoisWho@gmail.com",
	},
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		resp, err := json.Marshal(PersonData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(resp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func GetOneUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		fmt.Println(r.URL.Query().Get("id"))
		idOfUser, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			fmt.Println(err)
		}

		user, ok := PersonData[idOfUser]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
		}

		resp, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		_, err = w.Write(resp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		//TODO
	}
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		//TODO
	}
}
