package workers

import (
	"sync"
	"strings"
	"github.com/gorilla/websocket"
	"github.com/Traffic-Map-Team/traffic-map-tracker/vehicles"
)

type Worker struct {
	Source chan []vehicles.Vehicle
	Quit chan struct{}
}

type WorkerSlice struct {
	sync.Mutex
	workers []*Worker
}

func (slice *WorkerSlice) Push(w *Worker) {
	slice.Lock()
	defer slice.Unlock()

	slice.workers = append(slice.workers, w)
}

func (slice *WorkerSlice) Iter(routine func(*Worker)) {
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


func (slice *WorkerSlice) Broadcast(source chan []vehicles.Vehicle) {
	for msg := range source {
		slice.Iter(func (w *Worker) { w.Source <- msg })
	}
}

func (w *Worker) Start(conn *websocket.Conn, lines []string) {
	w.Source = make(chan []vehicles.Vehicle, 10)
	go func() {
		for {
			select {
			case vs := <- w.Source:
				var response []vehicles.Vehicle
				for _, v := range vs {
					if contains(lines, v.Line) || contains(lines, strings.ToLower(v.Line)) {
						response = append(response, v)
					}
				}
				err := conn.WriteJSON(response)
				if err != nil {
					return
				}
			case <- w.Quit:
				return
			}
		}
	}()
}
