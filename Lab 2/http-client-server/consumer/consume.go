package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main() {
	resp, _ := http.Get("http://localhost:8080/laureates") // http.Get("http://api.nobelprize.org/v1/laureate.json?bornCountryCode=mx")
	body, _ := ioutil.ReadAll(resp.Body)
	var mx Premiados
	json.Unmarshal(body, &mx)
	fmt.Println(mx)
}
