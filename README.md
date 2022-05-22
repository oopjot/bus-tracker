# Bus tracker
[Gdansk Open API](https://ckan.multimediagdansk.pl/dataset/tristar/resource/0683c92f-7241-4698-bbcc-e348ee355076).

## Running
### Try it out with Docker
```console
docker run -d -p 8000:8000 oopjot/tracker
```

## Usage
Information about lines X, Y and Z are served via websocket
```console
ws://0.0.0.0:8000/vehicles?lines=X,Y,Z
```

To test it, just reach the endpoint with ws tool (e. g. [wscat](https://github.com/websockets/wscat)).
```console
wscat -c "ws://0.0.0.0:8000/vehicles?lines=10,100,111"
```

## Data
Then, JSON is emmited.
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
B is a calculated bearing angle.

## Todo
- emit batch of data on connection
- move port number to some config file, or use env variable
- integrate other APIs

