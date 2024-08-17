package main

import (
	"HandlersTask/user_handlers"
	"database/sql"
	"fmt"
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

	connStr := "user=postgres password=postgres dbname=person_data sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from person_data")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var products []user_handlers.Person

	for rows.Next() {
		p := user_handlers.Person{}
		err := rows.Scan(&p.Id, &p.Name, &p.Age, &p.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Println(p.Id, p.Name, p.Age, p.Email)
	}
}
