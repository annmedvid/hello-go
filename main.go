package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", RootEndpoint)
	myRouter.HandleFunc("/ports", PortsEndpoint)
	myRouter.HandleFunc("/port", CreatePortEndpoint).Methods("POST")
	myRouter.HandleFunc("/port/{id}", UpdatePortEndpoint).Methods("PUT")
	myRouter.HandleFunc("/port/{id}", DeletePortEndpoint).Methods("DELETE")
	myRouter.HandleFunc("/port/{id}", PortEndpoint)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func applyPortsData() {
	fmt.Println("Open ports.json")
	jsonFile, _ := os.Open("ports.json")
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Ports)
}

func main() {
	fmt.Println("Ports REST API 2.0 - Mux Routers")
	applyPortsData()
	handleRequests()
}
