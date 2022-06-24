package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var store = make(map[string]interface{})

// Get returns the value of the key
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	_, ok := store[vars["key"]]
	if !ok {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(store[vars["key"]])
}

// Search returns the keys that match the prefix or suffix
func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var arr []string
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	queries := r.Form
	if len(queries["prefix"]) > 0 {
		for k := range store {
			if strings.HasPrefix(k, queries["prefix"][0]) {
				arr = append(arr, k)
			}
		}
	}
	if len(queries["suffix"]) > 0 {
		for k := range store {
			if strings.HasSuffix(k, queries["suffix"][0]) {
				arr = append(arr, k)
			}
		}
	}
	_ = json.NewEncoder(w).Encode(arr)
}

// Set sets the value of the key
func Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var new map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		panic(err)
	}
	for k, v := range new {
		store[k] = v
	}
	_ = json.NewEncoder(w).Encode(store)
}
