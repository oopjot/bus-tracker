# tracker
Tracker microservice

## Uruchamianie
### Docker
```console
docker run -d -p 8000:8000 oopjot/tracker
```
### Lokalnie
```console
go run main.go
```

## Użycie
Informacje o pojazdach linii X, Y oraz Z udostępniane są poprzez websockets pod adresem
```console
ws://0.0.0.0:8000/vehicles?lines=X,Y,Z
```
Aby przetestować działanie użyj narzędzia do websocketów, np. [wscat](https://github.com/websockets/wscat).
```console
wscat -c "ws://0.0.0.0/vehicles?linex=10,100,111"
```

## Dane
Emitowana jest lista obiektów Vehicle w JSONie, zgodnie z filtrami
```go
type Vehicle struct {
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
  B float64 `json:"b"`
}
```

## Todo
- przechować gdzieś i wysłać wiadomość od razu po połączeniu

