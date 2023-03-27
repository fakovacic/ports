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

- define interface Ports in parser service with only methods used by parser

- remove unused errors functions

- store folder can be extracted from ports - ports store implementation can be defined in additional /store/inmem/ports folder

- create models folder, adding only bussiness logic models ( without json or other tags ), additionaly adding models specific for each layer ( parsing and store with appropriate tags is needed ) 