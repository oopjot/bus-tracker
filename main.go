package main

import (
	"log"
	"net/http"
	"time"
	"github.com/Traffic-Map-Team/traffic-map-tracker/vehicles"
	"github.com/Traffic-Map-Team/traffic-map-tracker/workers"
	"github.com/Traffic-Map-Team/traffic-map-tracker/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var targetHash string
	sink := make(chan []vehicles.Vehicle)
	var workers workers.WorkerSlice
	var client = &http.Client{Timeout: 10 * time.Second}

	go vehicles.GetVehiclesData(client, &targetHash, sink)
	go workers.Broadcast(sink)

	r := mux.NewRouter()
	r.HandleFunc("/vehicles", handlers.VehiclesHandler(&workers))

	log.Fatal(http.ListenAndServe(":8000", r))

}
