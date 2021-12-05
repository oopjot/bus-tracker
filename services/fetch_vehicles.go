package services

import (
	"log"
	"net/http"
	"encoding/json"
	"time"
	"crypto/md5"
	"fmt"
	"errors"
	"math"
)



type Vehicle struct {
	VehicleResponse
	B float64
}

func (vr *VehiclesResponse) Find(id int) (VehicleResponse, error) {
	for _, vehicle := range vr.Vehicles {
		if vehicle.Id == id {
			return vehicle, nil
		}
	}
	return VehicleResponse{}, errors.New("Nie ma")
}

func getAllVehicles(client *http.Client, target *VehiclesResponse) error {

	r, err := client.Get("https://ckan2.multimediagdansk.pl/gpsPositions")
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(&target)
}

func GetVehiclesData(client *http.Client, targetHash *string, target chan []Vehicle) error {
	var stored [3]VehiclesResponse
	for {
		fmt.Println("------------------------------------------------------------")
		var data VehiclesResponse
		getAllVehicles(client, &data)
		log.Printf("Liczba pojazdów: %d\n", len(data.Vehicles))
		log.Printf("Ostatni updejt: %s\n", data.LastUpdate)
		h := md5.New()
		dataBytes := []byte(fmt.Sprintf("%v", data))
		h.Write(dataBytes)

		dataHashed := fmt.Sprintf("%x", h.Sum(nil))
		log.Println("Zahaszowane: ", dataHashed)

		if dataHashed == *targetHash {
			log.Println("To samo, nic nie robię")
		} else {
			log.Println("Podmieniam")
			targetHash = &dataHashed
			log.Println("Wpisuję do kanału")
			stored[2] = stored[1]
			stored[1] = stored[0]
			stored[0] = data
			p := process(stored)
			target <- p
		}

		time.Sleep(3 * time.Second)
	}
}



func getBearingAngle(latA, lonA, latB, lonB float64) float64 {
	delta := latB - latA
	X := math.Cos(lonA) * math.Sin(lonB) - math.Sin(lonA) * math.Cos(lonB) * math.Cos(delta)
	Y := math.Cos(lonA) * math.Sin(lonB) - math.Sin(lonA) * math.Cos(lonB) * math.Cos(delta)
	return math.Atan2(X, Y) * 57.29
}

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
		v := Vehicle{
		}
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

