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
  DataGenerated string
  Line string
  Route string
  VehicleCode string
  VehicleService string
  Id int
  Speed: int 
  Delay int
  Lat float64
  Lon float64
  GpsQuality int
  B float64
}
```
B jest kątem obrotu pojazdu

## Todo
- przechować gdzieś i wysłać wiadomość od razu po połączeniu

