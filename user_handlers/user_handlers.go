package user_handlers

import (
	"HandlersTask/pkg"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Person struct {
	Id    int    `json:"id" db:"id" `
	Name  string `json:"name" db:"name"`
	Age   int    `json:"age" db:"age"`
	Email string `json:"email," db:"email"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		pkg.AddHeaders(w)
		db, err := pkg.ConnectToDB()
		if err != nil {
			log.Fatal(err)
		}

		rows, err := db.Query("select * from users")
		if err != nil {
			panic(err)
		}

		defer func(rows *sql.Rows) {
			if err := rows.Close(); err != nil {
				panic(err)
			}
		}(rows)

		var user = []Person{}
		for rows.Next() {
			u := Person{}
			err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.Email)
			if err != nil {
				panic(err)

			}
			user = append(user, u)
		}

		resp, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			pkg.ErrorResponse(w, "Service error", "Ошибка на сервере", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(resp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			pkg.ErrorResponse(w, "Service error", "Ошибка на сервере", http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		pkg.AddHeaders(w)

		db, err := pkg.ConnectToDB()
		if err != nil {
			log.Fatal(err)
		}

		row := db.QueryRow("select * from users where id = $1", r.URL.Query().Get("id"))

		var user Person
		err = row.Scan(&user.Id, &user.Name, &user.Age, &user.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				pkg.ErrorResponse(w, "No rows", "Такого пользователя нет", http.StatusInternalServerError)
				return
			}
		}

		bytesBody, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}

		_, err = w.Write(bytesBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		pkg.AddHeaders(w)

		db, err := pkg.ConnectToDB()
		if err != nil {
			log.Fatal(err)
		}

		var s Person
		err = json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			pkg.ErrorResponse(w, "Not found", "Ошибка декода, возможно типы данных", http.StatusNotFound)
		}
		err = r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			pkg.ErrorResponse(w, "Body close error", "Ошибка на сервере", http.StatusInternalServerError)
		}

		res, err := db.Exec("insert into users (name, age, email) VALUES ($1,$2,$3) returning id", s.Name, s.Age, s.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		(res.LastInsertId()) //последний полученный айдишник, но библа pq выдает ошибку "LastInsertId is not supported by this driver"

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		pkg.AddHeaders(w)

		db, err := pkg.ConnectToDB()
		if err != nil {
			log.Fatal(err)
		}

		res, err := db.Exec("delete from users where id = $1", r.URL.Query().Get("id"))
		if err != nil {
			return
		}
		(res.RowsAffected()) //драйвер не умеет в это

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
