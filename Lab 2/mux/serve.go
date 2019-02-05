package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

var (
	items map[string]Item
)

func getItem(w http.ResponseWriter, r *http.Request) {
	handleCors(w, r)
	vars := mux.Vars(r)
	id := vars["id"]
	for ID, item := range items {
		if ID == id {
			var marshalled []byte
			var err error
			marshalled, err = json.Marshal(item)
			if err != nil {
				fmt.Fprintf(w, "Error!")
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.Write(marshalled)
			}
		}
	}
}

func getItems(w http.ResponseWriter, r *http.Request) {
	handleCors(w, r)
	w.Header().Set("Content-Type", "application/json")
	var itemArray []Item = []Item{}
	for _, v := range items {
		itemArray = append(itemArray, v)
	}
	marshalled, err := json.Marshal(itemArray)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshalled)
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	handleCors(w, r)
	vars := mux.Vars(r)
	id := vars["id"]
	for ID, item := range items {
		if ID == id {
			marshalled, err := json.Marshal(item)
			if err != nil {
				fmt.Fprintf(w, "Error!")
			} else {
				delete(items, item.ID)
				w.Header().Set("Content-Type", "application/json")
				w.Write(marshalled)
			}
		}
	}
}

func deleteItems(w http.ResponseWriter, r *http.Request) {
	handleCors(w, r)
	var itemArray []Item = []Item{}
	for _, v := range items {
		itemArray = append(itemArray, v)
	}
	marshalled, err := json.Marshal(itemArray)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	} else {
		items = make(map[string]Item)
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshalled)
	}
}

func editItem(w http.ResponseWriter, r *http.Request) {
	handleCors(w, r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item Item
	json.Unmarshal(body, &item)
	if items[item.ID] == (Item{}) {
		fmt.Fprintf(w, "Error!")
	} else {
		items[item.ID] = item
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

func addItem(w http.ResponseWriter, r *http.Request) {
	handleCors(w, r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}
	var item Item
	json.Unmarshal(body, &item)
	if items[item.ID] == (Item{}) {
		items[item.ID] = item
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	} else {
		fmt.Fprintf(w, "Error!")
	}
}

func handleCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	w.Header().Set("Vary", "Origin")
	w.Header().Set("Vary", "Access-Control-Request-Method")
	w.Header().Set("Vary", "Access-Control-Request-Headers")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		return
	}
}

func main() {
	items = make(map[string]Item)
	r := mux.NewRouter()
	r.HandleFunc("/", handleCors).Methods("OPTIONS")
	r.HandleFunc("/item/", getItems).Methods("GET")
	r.HandleFunc("/item/", handleCors).Methods("OPTIONS")
	r.HandleFunc("/item/{id}", getItem).Methods("GET")
	r.HandleFunc("/item/{id}", handleCors).Methods("OPTIONS")
	r.HandleFunc("/item/", addItem).Methods("POST")
	r.HandleFunc("/item/", editItem).Methods("PUT")
	r.HandleFunc("/item/{id}", deleteItem).Methods("DELETE")
	r.HandleFunc("/item/", deleteItems).Methods("DELETE")

	http.Handle("/", r)
	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		log.Fatal("error en el servidor : ", err)
		return
	}
}
