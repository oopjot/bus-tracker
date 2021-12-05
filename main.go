package main

import (
	"log"
	"net/http"
	"time"
	"github.com/Traffic-Map-Team/traffic-map-tracker/services"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"sync"
	"strings"
)

type worker struct {
	source chan []services.Vehicle
	quit chan struct{}
}

type workerSlice struct {
	sync.Mutex
	workers []*worker
}

func (slice *workerSlice) Push(w *worker) {
	slice.Lock()
	defer slice.Unlock()

	slice.workers = append(slice.workers, w)
}

func (slice *workerSlice) Iter(routine func(*worker)) {
	slice.Lock()
	defer slice.Unlock()

	for _, worker := range slice.workers {
		routine(worker)
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (w *worker) Start(conn *websocket.Conn, lines []string) {
	w.source = make(chan []services.Vehicle, 10)
	go func() {
		for {
			select {
			case vehicles := <- w.source:
				var response []services.Vehicle
				for _, v := range vehicles {
					if contains(lines, v.Line) || contains(lines, strings.ToLower(v.Line)) {
						response = append(response, v)
					}
				}
				err := conn.WriteJSON(response)
				if err != nil {
					fmt.Println("Zerwano połączenie, kończę")
					return
				}
			case <- w.quit:
				fmt.Println("Kończę")
				return
			}
		}
	}()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}


func VehiclesHandler(workers *workerSlice) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		lines := req.URL.Query().Get("lines")
		linesArr := strings.Split(lines, ",")
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		worker := &worker{}
		worker.quit = make(chan struct{})

		ws, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Println(err)
			var q struct{}
			worker.quit <- q
		}
		worker.Start(ws, linesArr)
		workers.Push(worker)
	}
}


func main() {
	var targetHash string
	vehicles := make(chan []services.Vehicle)
	var client = &http.Client{Timeout: 10 * time.Second}
	var workers workerSlice
	go services.GetVehiclesData(client, &targetHash, vehicles)

	go func() {
		for msg := range vehicles {
			workers.Iter(func (w *worker) { w.source <- msg })
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/vehicles", VehiclesHandler(&workers))

	log.Fatal(http.ListenAndServe(":8000", r))

}
