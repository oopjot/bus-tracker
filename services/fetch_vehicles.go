package services

import (
	"log"
	"net/http"
	"encoding/json"
	"time"
	"crypto/md5"
	"fmt"
)

func getAllVehicles(client *http.Client, target *VehiclesResponse) error {

	r, err := client.Get("https://ckan2.multimediagdansk.pl/gpsPositions")
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(&target)
}

func GetVehiclesData(client *http.Client, target *VehiclesResponse, targetHash *string) error {
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
			target = &data
			targetHash = &dataHashed
		}

		time.Sleep(3 * time.Second)
	}
}

