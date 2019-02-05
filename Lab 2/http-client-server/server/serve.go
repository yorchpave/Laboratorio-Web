package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Premiados struct {
	Ganadores []Ganador `json:"laureates"`
}

type Ganador struct {
	Id      string  `json:"id"`
	Name    string  `json:"firstname"`
	Surname string  `json:"surname"`
	Prizes  []Prize `json:"prizes"`
}

type Prize struct {
	Category string `json:"category"`
}

func getLaureates(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var laureates Premiados = Premiados{
			Ganadores: []Ganador{
				Ganador{
					Id:      "282",
					Name:    "Mario J.",
					Surname: "Molina",
					Prizes:  []Prize{Prize{Category: "Chemistry"}},
				},
			},
		}
		var marshalled []byte
		var err error
		marshalled, err = json.Marshal(laureates)
		if err != nil {
			fmt.Fprintf(w, "Error!")
		} else {
			w.Header().Set("Content-Type", "application/json")
			fmt.Println(laureates)
			w.Write(marshalled)
		}
	}
}

func main() {
	http.HandleFunc("/laureates", getLaureates)
	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		log.Fatal("error en el servidor : ", err)
		return
	}
}
