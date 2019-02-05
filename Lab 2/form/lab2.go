package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func inspectRequest(response http.ResponseWriter, request *http.Request) {
	fmt.Println("URL: ", request.URL, "\nHeader: ", request.Header)
	fmt.Fprintf(response, "URL: ", request.URL, "\nHeader: ", request.Header)
}

func formHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		http.ServeFile(response, request, "form.html")
	}
	if request.Method == "POST" {
		request.ParseForm()
		response.Header().Set("Content-Type", "application/json")
		fmt.Println(request.Form)
		json.NewEncoder(response).Encode(request.Form)
	}
}

func main() {
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/inspectRequest", inspectRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
