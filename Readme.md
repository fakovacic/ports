# Ports service

## Run

```
make compile

docker-compose up
```

```
# TODO

- id/unique key not clear - should create uuid or similar for new, or use map key from ports.json file
( seems same as unlocs field, but defined as slice ) or maybe code from struct, did not check if unique

- add log wrapper - currently using zerolog struct directly
