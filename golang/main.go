package main

import (
	"net/http"
)

func coastersHandler(w http.ResponseWriter, r *http.Request) {

}
func main() {
	http.HandleFunc("/coasters", coastersHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
