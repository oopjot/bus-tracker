package services

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

