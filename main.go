package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	  "io/ioutil"
)

type Player struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

var players []Player

func initPlayers(){

	players = append(players, Player{"1", "Ginny", 1})
	players = append(players, Player{"2", "Tim", 35})
	players = append(players, Player{"3", "Jack", 3})
	players = append(players, Player{"4", "Porter", 6})
	players = append(players, Player{"5", "Steph", 35})
}

func main() {
	initPlayers()

	r := mux.NewRouter()
	r.HandleFunc("/players", getPlayersHandler)
	r.HandleFunc("/players/{id}", putPlayersHandler).Methods("PUT")
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


func getPlayersHandler(w http.ResponseWriter, r *http.Request) {
	
	playersJson, err := json.Marshal(players)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(playersJson)
}

func putPlayersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id )
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	var player Player
	err = json.Unmarshal(body, &player)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	//need to check this index
	players[i-1] = player
	log.Printf("Updated Players: %#v\n",players)

	playersJson, err := json.Marshal(player)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(playersJson)
}