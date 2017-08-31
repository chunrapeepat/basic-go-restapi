package main

import (
  "fmt"
  "log"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
)

type Person struct {
  ID string `json:"id,omitempty"`
  Firstname string `json:"firstname,omitempty"`
  Lastname string `json:"lastname,omitempty"`
  Address *Address `json:"address,omitempty"`
}

type Address struct {
  City string `json:"city,omitempty"`
  State string `json:"state,omitempty"`
}

var people []Person

func getPersonEndpoint(w http.ResponseWriter, req *http.Request) {
  json.NewEncoder(w).Encode(people)
}

func getPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for _, item := range people {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Person{})
}

func main() {
  router := mux.NewRouter()

  people = append(people, Person{ ID: "1", Firstname: "Nic", Lastname: "Raboy", Address: &Address{ City: "Bangkapi", State: "Bangkok" } })
  people = append(people, Person{ ID: "2", Firstname: "Chun", Lastname: "Rapeepat", Address: &Address{ City: "Bangkapi", State: "Changmai" } })

  router.HandleFunc("/person", getPersonEndpoint).Methods("GET")
  router.HandleFunc("/person/{id}", getPeopleEndpoint).Methods("GET")
  log.Fatal(http.ListenAndServe(":3000", router))

  fmt.Println("> Server is running on port 3000")
}
