package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int8 `json:"age"`
	Class []int `json:"class"`
}

type Students []Student

func main()  {
	fmt.Println("My Website is running...")

	http.HandleFunc("/", MyHomePage)
	http.HandleFunc("/about", MyAboutPage)
	http.HandleFunc("/api/music", MyMusicApi)
	http.HandleFunc("/api/student", MyAStudentApi)
	http.HandleFunc("/api/students", MyListStudentApi)
	log.Fatal(http.ListenAndServe(":333", nil))
}

func MyAStudentApi(w http.ResponseWriter, r *http.Request)  {
	var student = Student{1, "Diep", 18, []int{1,2,3}}
	json.NewEncoder(w).Encode(student)
}

func MyListStudentApi(w http.ResponseWriter, r *http.Request)  {
	var listStudent = Students{
		Student{1, "Diep", 18, []int{1,2,3}},
		Student{2, "Nam", 18, []int{1,2,3}},
		Student{3, "Trong", 18, []int{1,2,3}},
		Student{4, "Tuan", 18, []int{1,2,3}},
	}
	json.NewEncoder(w).Encode(listStudent)
}

func MyMusicApi(w http.ResponseWriter, r *http.Request)  {
	var data = map[string]interface{}{
		"name" : "Loi Xin Loi Cua Mot Dan Choi",
		"casi" : "Duy Manh",
	}

	json.NewEncoder(w).Encode(data)
}

func MyHomePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "<h1>This is home page</h1>")
}

func MyAboutPage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "<h1>This is about page</h1>")
}