# Pack management service

This is a simply application which enables us to insert package types and calculate efficient order shippings.

## Running

> To Run this project, you will need to have docker and Go 1.21.
> For brevity, this Project does not set any DB container volume, ***so each time you stop the database you will need to re-run the migrations***

### 1 - spin up the database

```sh
make db-up
```

### 2 - run the migrations

```sh
make migration-up
```

### 3 - run the api

```sh
make api
```

## Trying out

There's a JSON file under the `/postman` directory with the collection containing all the requests

## Testing

This project has a single unit test covering the core calculation, you can run it using:

```sh
make test
```
