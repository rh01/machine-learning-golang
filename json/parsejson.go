package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const citiBenkenycURL = "http://gbfs.citibikenyc.com/gbfs/en/station_status.json"

type stationData struct {
	LastUpdated int `json:last_updated`
	TTL         int `json:ttl`
	Data        struct {
		Stations []station `json: stations`
	}
}

type station struct {
	StationId              string `json: station_id`
	NumBikesAvailable      int    `json:num_bikes_available`
	NumEbikesAvailable     int    `json:num_ebikes_available`
	NumBikesDisabled       int    `json:num_bikes_disabled`
	NumDocksAvailable      int    `json: num_docks_available`
	NumDocksDisabled       int    `json:num_docks_disabled`
	IsInstalled            int    `json:is_installed`
	IsRenting              int    `json:is_renting`
	IsReturning            int    `json:is_returning`
	LastReported           int    `json:last_reported`
	EightdHasAvailableKeys bool   `json:eightd_has_available_keys`
}

func main() {
	r, err := http.Get(citiBenkenycURL)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var sd stationData

	if err := json.Unmarshal(b, &sd); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n\n", sd.Data.Stations[0])

	outputData, err := json.Marshal(sd)
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}

	// 保存
	if err := ioutil.WriteFile("./data/citibike.json", outputData, 0644); err != nil {
		log.Fatal(err)
	}

}
