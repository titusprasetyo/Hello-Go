package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "model"
)

var people []model.Person

func main(){
	people = append(people, model.Person{ID: "1", Firstname: "John", Lastname: "Doe"})
    router := mux.NewRouter()
    router.HandleFunc("/people", GetPeople).Methods("GET")
    log.Fatal("EROT : ", http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

