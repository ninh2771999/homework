package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const jsonContentType = "application/json"

var (
	listToken = map[int]string{1: "tk1", 2: "tk2"}
	students  []Student
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func contentTypeCheckingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		reqAuthorization := r.Header.Get("Authorization")

		if reqContentType != jsonContentType {
			fmt.Println("request not allow")
			json.NewEncoder(w).Encode(http.StatusBadRequest)
			json.NewEncoder(w).Encode("request only allow Content-Type: application/json")
			return
		}
		check := false
		for _, x := range listToken {
			if x == reqAuthorization {
				check = true
				break
			}
		}
		if !check {
			fmt.Println("not login")
			json.NewEncoder(w).Encode(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Please login")
			return
		}
		fmt.Println(" request is allow")
		next.ServeHTTP(w, r)
	})
}
func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("access welcome")
	json.NewEncoder(w).Encode(http.StatusOK)
	json.NewEncoder(w).Encode("Welcome to our website")
}
func getStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("access get list students")
	json.NewEncoder(w).Encode(http.StatusOK)
	json.NewEncoder(w).Encode(students)

}

func main() {
	students = append(students, Student{Id: 1, Name: "HS A"})
	students = append(students, Student{Id: 2, Name: "HS B"})
	students = append(students, Student{Id: 3, Name: "HS C"})

	router := mux.NewRouter().StrictSlash(true)
	router.Use(contentTypeCheckingMiddleware)
	router.Methods(http.MethodGet).Path("/").HandlerFunc(welcome)
	router.Methods(http.MethodGet).Path("/students").HandlerFunc(getStudents)

	err := http.ListenAndServe(":8888", router)
	if err != nil {
		panic(err)
	}
}
