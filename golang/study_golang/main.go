package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type student struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Class string  `json:"class"`
  
}


var students = []student{
    {ID: "1", Name: "Ninh", Class: "12A9"},
	{ID: "1", Name: "tien", Class: "12A7"},
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/student", studentHandler).Methods("GET")
	r.HandleFunc("/goodbye", goodbyeHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
	r.Use(middleware)
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello i am ninh")
	c.IndentedJSON(http.StatusOK, albums)
	
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware")
		next.ServeHTTP(w, r)
	})
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Goodbye guys !")
}
