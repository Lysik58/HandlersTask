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
			w.WriteHeader(http.StatusInternalServerError)
			ErrorResponse(w, "Service Error", "Ошибка на сервере", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(resp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			ErrorResponse(w, "Service Error", "Ошибка на сервере", http.StatusInternalServerError)

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
			w.WriteHeader(http.StatusInternalServerError)
			ErrorResponse(w, "service error", "Ошибка на сервере", http.StatusInternalServerError)
		}

		user, ok := PersonData[idOfUser]
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			ErrorResponse(w, "user not found", "Такого человека нет в базе", http.StatusNotFound)
			return
		}

		resp, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ErrorResponse(w, "Marshall error", "Ошибка на сервере", http.StatusInternalServerError)
		}

		_, err = w.Write(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ErrorResponse(w, "Response error", "Ошибка на сервере", http.StatusInternalServerError)
		}
	}
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		var s Person
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ErrorResponse(w, "not found", "Ошибка декода, возможно типы данных", http.StatusNotFound)
		}
		err = r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ErrorResponse(w, "Body close error", "Ошибка на сервере", http.StatusInternalServerError)
		}
		PersonData[len(PersonData)+1] = Person{
			Id:    len(PersonData) + 1,
			Name:  s.Name,
			Age:   s.Age,
			Email: s.Email,
		}
		fmt.Println(s)
	}

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		idOfUser, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ErrorResponse(w, "service error", "Ошибка на сервере", http.StatusInternalServerError)
		}

		delete(PersonData, idOfUser)
		_, err = w.Write([]byte(strconv.Itoa(len(PersonData))))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ErrorResponse(w, "delete error", "Ошибка на сервере", http.StatusInternalServerError)
		}
	}
}
