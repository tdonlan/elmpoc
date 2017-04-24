package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/players", getPlayersHandler)
	r.HandleFunc("/api/hello", helloHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("frontend")))

	http.Handle("/", r)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World! %#v", r)
}

type Player struct {
	Id    string    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func getPlayersHandler(w http.ResponseWriter, r *http.Request) {
	players := []Player{}

	players = append(players, Player{"1", "Ginny", 1})
	players = append(players, Player{"2", "Tim", 35})
	players = append(players, Player{"3", "Jack", 3})
	players = append(players, Player{"4", "Porter", 6})
	players = append(players, Player{"5", "Steph", 35})

	playersJson, err := json.Marshal(players)
	if err != nil {
		w.WriteHeader(500)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(playersJson)
}
