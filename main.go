package main

import (
    "fmt"
    "log"
    "io"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Port struct {
    Code string `json:"code"`
    Name string `json:"name"`
    City string `json:"city"`
    Country string `json:"country"`
}

var Ports []Port

func portsEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: portsEndpoint")
    json.NewEncoder(w).Encode(Ports)
}

func portEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: portEndpoint")
    vars := mux.Vars(r)
    portCode := vars["code"]

    for _, port := range Ports {
        if port.Code == portCode {
            json.NewEncoder(w).Encode(port)
        }
    }
}

func createPortEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: createPortEndpoint")
    reqBody, _ := io.ReadAll(r.Body)
    var port Port
    json.Unmarshal(reqBody, &port)
    Ports = append(Ports, port)
    json.NewEncoder(w).Encode(port)
}

func deletePortEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: deletePortEndpoint")
    vars := mux.Vars(r)
    portCode := vars["code"]

    for index, port := range Ports {
        if port.Code == portCode {
            Ports = append(Ports[:index], Ports[index+1:]...)
        }
    }
}

func rootEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Simple REST API for ports handling. Welcome.")
    fmt.Println("Endpoint Hit: rootEndpoint")
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", rootEndpoint)
    myRouter.HandleFunc("/ports", portsEndpoint)
    myRouter.HandleFunc("/port", createPortEndpoint).Methods("POST")
    myRouter.HandleFunc("/port/{code}", deletePortEndpoint).Methods("DELETE")
    myRouter.HandleFunc("/port/{code}", portEndpoint)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    fmt.Println("Ports REST API 2.0 - Mux Routers")
    Ports = []Port{
        Port{Code: "35700", Name: "Goya", City: "Goya", Country: "Argentina"},
        Port{Code: "60237", Name: "Melbourne", City: "Melbourne", Country: "Australia"},
    }
    handleRequests()
}
