# Ports service

## Run

```
make compile
make docker-build
docker-compose up
```

## Routes

### POST - /ports

- create ports

#### Example requests

```
curl -X POST http://localhost:8080/ports -d '{"port":{"name":"Rijeka"}}'
```

### PUT - /ports/:id

- update ports

#### Example requests

```
curl -X PUT http://localhost:8080/ports/HRRJK -d '{"port":{"name":"Rijeka", "city":"Rijeka"}}'
```
# TODO

- id/unique key not clear - should create uuid or similar for new, or use map key from ports.json file
( seems same as unlocs field, but defined as slice ) or maybe code from struct, did not check if unique

- add log wrapper - currently using zerolog struct directly

## parser

- I would possibly use json.NewDecoder or some external package to parse json file and store data in "database", but need to investigate more about it