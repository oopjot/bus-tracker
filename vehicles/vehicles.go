package vehicles

import (
	"log"
	"encoding/json"
	"net/http"
	"crypto/md5"
	"math"
	"errors"
	"fmt"
	"time"
)

type VehicleResponse struct {
	DataGenerated string `json:"DataGenerated"`
	Line string `json:"Line"`
	Route string `json:"Route"`
	VehicleCode string `json:"VehicleCode"`
	VehicleService string `json:"VehicleService"`
	Id int `json:"VehicleId"`
	Speed int `json:"Speed"`
	Delay int `json:"Delay"`
	Lat float64 `json:"Lat"`
	Lon float64 `json:"Lon"`
	GpsQuality int `json:"GPSQuality"`
}

type VehiclesResponse struct {
	LastUpdate string `json:"LastUpdateData"`
	Vehicles []VehicleResponse `json:"Vehicles"`
}

type Vehicle struct {
	VehicleResponse
	B float64 `json:"b"`
}

// find vehicle by Id
func (vr *VehiclesResponse) Find(id int) (VehicleResponse, error) {
	for _, vehicle := range vr.Vehicles {
		if vehicle.Id == id {
			return vehicle, nil
		}
	}
	return VehicleResponse{}, errors.New(fmt.Sprintf("Vehicle not found: %d", id))
}

// fetch vehicles from public api
func getAllVehicles(client *http.Client, target *VehiclesResponse) error {

	r, err := client.Get("https://ckan2.multimediagdansk.pl/gpsPositions")
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(&target)
}

// utilize fetch function
func GetVehiclesData(client *http.Client, targetHash *string, target chan []Vehicle) error {
	var stored [3]VehiclesResponse
	for {
		var data VehiclesResponse
		getAllVehicles(client, &data)
		h := md5.New()
		dataBytes := []byte(fmt.Sprintf("%v", data))
		h.Write(dataBytes)

		dataHashed := fmt.Sprintf("%x", h.Sum(nil))

		if dataHashed != *targetHash {
			targetHash = &dataHashed
			log.Println("New data: " + dataHashed + " at " + data.LastUpdate)
			stored[2] = stored[1]
			stored[1] = stored[0]
			stored[0] = data
			p := process(stored)
			target <- p
		}
		time.Sleep(3 * time.Second)
	}
}


// bearing angle helper func
func getBearingAngle(latA, lonA, latB, lonB float64) float64 {
	delta := lonB - lonA
	X := math.Cos(latB) * math.Sin(delta)
	Y := math.Cos(latA) * math.Sin(latB) - math.Sin(latA) * math.Cos(latB) * math.Cos(delta)
	return math.Atan2(X, Y) * 57.29
}

// process fetched vehicles
func process(data [3]VehiclesResponse) (result []Vehicle) {
	latest := data[0]
	prev := data[1]
	for _, vl := range latest.Vehicles {
		vp, err := prev.Find(vl.Id)
		if err != nil {
			vp, err = data[2].Find(vl.Id)
			if err != nil {
				log.Println(err)
			}
		}
		B := getBearingAngle(vp.Lat, vp.Lon, vl.Lat, vl.Lon)
		v := Vehicle{}
		v.DataGenerated = vl.DataGenerated
		v.Line = vl.Line
		v.Route = vl.Route
		v.VehicleCode = vl.VehicleCode
		v.VehicleService = vl.VehicleService
		v.Id = vl.Id
		v.Speed = vl.Speed
		v.Delay = vl.Delay
		v.Lat = vl.Lat
		v.Lon = vl.Lon
		v.GpsQuality = vl.GpsQuality

		v.B = B
		result = append(result, v)
	}
	return
}

