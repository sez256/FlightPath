package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/calculate", getFlightPath)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getFlightPath(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var flightMap map[string]string
	err := json.NewDecoder(r.Body).Decode(&flightMap)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	flightList := make(map[string]uint32)

	var start, end string
	for k, v := range flightMap {
		flightList[k] += 1
		flightList[v] += 2
	}

	for k, v := range flightMap {
		if flightList[k] == 1 {
			start = k
		}

		if flightList[v] == 2 {
			end = v
		}
	}

	// // Construct the flight path, this makes the full path of flight
	// flightPath := []string{start}
	// for {
	// 	v, ok := flightMap[start]
	// 	fmt.Println(start, v, ok)
	// 	if !ok {
	// 		break
	// 	}
	// 	flightPath = append(flightPath, v)
	// 	start = v
	// }

	type FlightPath struct {
		Start string `json:"start"`
		End   string `json:"end"`
	}

	flightPath := FlightPath{Start: start, End: end}

	response, err := json.Marshal(flightPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
