package wins

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// ListHandler handles requests to list all existing Wins
func ListHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[Wins] List")
	wins, err := list()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wins)
}

// CreateHandler handles requests to create a new Win
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[Wins] Create")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	win, err := create(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(win)
}
