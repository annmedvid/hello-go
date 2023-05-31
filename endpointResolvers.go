package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"example/user/hello-go/utils"

	"github.com/gorilla/mux"
)

func PortsEndpoint(w http.ResponseWriter, r *http.Request) {
	utils.Logger.Info("Endpoint Hit: portsEndpoint")
	json.NewEncoder(w).Encode(Ports)
}

func PortEndpoint(w http.ResponseWriter, r *http.Request) {
	utils.Logger.Info("Endpoint Hit: portEndpoint")
	vars := mux.Vars(r)
	portId := vars["id"]

	json.NewEncoder(w).Encode(Ports[portId])
}

func CreatePortEndpoint(w http.ResponseWriter, r *http.Request) {
	utils.Logger.Info("Endpoint Hit: createPortEndpoint")
	reqBody, _ := io.ReadAll(r.Body)
	var input PortInput
	json.Unmarshal(reqBody, &input)
	Ports[input.Id] = input.Port
	json.NewEncoder(w).Encode(input.Port)
}

func UpdatePortEndpoint(w http.ResponseWriter, r *http.Request) {
	utils.Logger.Info("Endpoint Hit: updatePortEndpoint")
	vars := mux.Vars(r)
	portId := vars["id"]

	reqBody, _ := io.ReadAll(r.Body)
	var updatedPort Port
	json.Unmarshal(reqBody, &updatedPort)

	Ports[portId] = updatedPort
	json.NewEncoder(w).Encode(updatedPort)
}

func DeletePortEndpoint(w http.ResponseWriter, r *http.Request) {
	utils.Logger.Info("Endpoint Hit: deletePortEndpoint")
	vars := mux.Vars(r)
	portId := vars["id"]

	delete(Ports, portId)
	json.NewEncoder(w).Encode(Ports)
}

func RootEndpoint(w http.ResponseWriter, r *http.Request) {
	utils.Logger.Info("Endpoint Hit: rootEndpoint")
	fmt.Fprintf(w, "Simple REST API for ports handling. Welcome.")
}
