package main

import (
	"log"
	"net/http"
	//"time"
	//"github.com/Traffic-Map-Team/traffic-map-tracker/services"
	//"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"strings"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}


func VehiclesHandler(w http.ResponseWriter, req *http.Request) {
	lines := req.URL.Query().Get("lines")

	fmt.Println(lines)
	linesArr := strings.Split(lines, ",")
	for _, line := range linesArr {
		fmt.Println(line)
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	_, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("New connection")

}

func main() {
	//var targetHash string
  //vehicles := make(chan []services.Vehicle)
	//var client = &http.Client{Timeout: 10 * time.Second}
  //go services.GetVehiclesData(client, vehicles, &targetHash, &date)
	r := mux.NewRouter()
	r.HandleFunc("/vehicles", VehiclesHandler)

	log.Fatal(http.ListenAndServe(":8000", r))

}
