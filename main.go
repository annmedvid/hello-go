package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

type Port struct {
    Name string `json:"name"`
    City string `json:"city"`
    Country string `json:"country"`
}

var Ports []Port

func portsEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: portsEndpoint")
    json.NewEncoder(w).Encode(Ports)
}

func rootEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Simple REST API for ports handling. Welcome.")
    fmt.Println("Endpoint Hit: rootEndpoint")
}

func handleRequests() {
    http.HandleFunc("/", rootEndpoint)
    http.HandleFunc("/ports", portsEndpoint)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    Ports = []Port{
        Port{Name: "Goya", City: "Goya", Country: "Argentina"},
        Port{Name: "Melbourne", City: "Melbourne", Country: "Australia"},
    }
    handleRequests()
}
