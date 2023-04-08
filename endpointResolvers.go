package main

import (
    "fmt"
    "io"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func PortsEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: portsEndpoint")
    json.NewEncoder(w).Encode(Ports)
}

func PortEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: portEndpoint")
    vars := mux.Vars(r)
    portCode := vars["code"]

    for _, port := range Ports {
        if port.Code == portCode {
            json.NewEncoder(w).Encode(port)
        }
    }
}

func CreatePortEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: createPortEndpoint")
    reqBody, _ := io.ReadAll(r.Body)
    var port Port
    json.Unmarshal(reqBody, &port)
    Ports = append(Ports, port)
    json.NewEncoder(w).Encode(port)
}

func UpdatePortEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: updatePortEndpoint")
    vars := mux.Vars(r)
    portCode := vars["code"]

    reqBody, _ := io.ReadAll(r.Body)
    var updatedPort Port
    json.Unmarshal(reqBody, &updatedPort)

    for index, port := range Ports {
        if port.Code == portCode {
            port.Name = updatedPort.Name
		    port.City = updatedPort.City
		    port.Country = updatedPort.Country
		    Ports[index] = port
		    json.NewEncoder(w).Encode(port)
        }
    }
}

func DeletePortEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: deletePortEndpoint")
    vars := mux.Vars(r)
    portCode := vars["code"]

    for index, port := range Ports {
        if port.Code == portCode {
            Ports = append(Ports[:index], Ports[index+1:]...)
        }
    }
}

func RootEndpoint(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Simple REST API for ports handling. Welcome.")
    fmt.Println("Endpoint Hit: rootEndpoint")
}
