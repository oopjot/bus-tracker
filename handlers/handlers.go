package handlers

import (
	"log"
	"strings"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/Traffic-Map-Team/traffic-map-tracker/workers"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func VehiclesHandler(slice *workers.WorkerSlice) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		lines := req.URL.Query().Get("lines")
		linesArr := strings.Split(lines, ",")
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		worker := &workers.Worker{}
		worker.Quit = make(chan struct{})

		ws, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Println(err)
			var q struct{}
			worker.Quit <- q
		}
		worker.Start(ws, linesArr)
		slice.Push(worker)
	}
}

