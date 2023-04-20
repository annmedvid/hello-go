package main

type Port struct {
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	City        string     `json:"city"`
	Country     string     `json:"country"`
	Alias       []string   `json:"alias"`
	Regions     []string   `json:"regions"`
	Coordinates [2]float64 `json:"coordinates"`
	Province    string     `json:"province"`
	Timezone    string     `json:"timezone"`
	Unlocs      []string   `json:"unlocs"`
}

type PortInput struct {
	Id   string `json:"id"`
	Port Port   `json:"port"`
}

var Ports map[string]Port
