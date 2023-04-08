package main

type Port struct {
    Code string `json:"code"`
    Name string `json:"name"`
    City string `json:"city"`
    Country string `json:"country"`
}

var Ports []Port
