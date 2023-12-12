# Pack management service

This is a simply application which enables us to insert package types and calculate efficient order shippings.

## Running

### Dependencies

- Docker

> For brevity, this Project does not set any DB container volume, ***so each time you stop the database you will need to re-run the migrations***

Just run:

```sh
docker compose up
```

under the hoods, it will spin up 3 containers.
1 - Database
2 - Migrator (short lived)
3 - API service

> the goose (migrator) image it's not optmized, so it will take some time to download all the dependencies

## Trying out

There's a JSON file under the `/postman` directory with the collection containing all the requests

## Testing

```sh
make test
```
