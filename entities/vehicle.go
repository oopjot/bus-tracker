type vehicleApiResponse struct {
	dataGenerated: string `json:"DataGenerated"`
	line: string `json:"Line"`
	route: string `json:"Route"`
	vehicleCode: string `json:"VehicleCode"`
	vehicleService: string `json:"VehicleService"`
	id: int `json:"VehicleId"`
	speed: int `json:"Speed"`
	delay: int `json:"Delay"`
	lat: float64 `json:"Lat"`
	lon: float64 `json:"Lon"`
	gpsQuality: int `json:"GPSQuality"`
}

type PublicApiResponse struct {
	dataGenerated: string `json:"DataGenerated"`
	vehicles: []vehicleApiResponse 'json:"Vehicles"'
}
