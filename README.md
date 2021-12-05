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

## Todo
- przechować gdzieś i wysłać wiadomość od razu po połączeniu

