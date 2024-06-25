package UserHandlers

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func PersonWriter(w http.ResponseWriter, r *http.Request) {
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
