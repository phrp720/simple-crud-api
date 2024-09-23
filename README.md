# RESTful CRUD JSON API  microservice


## Description

A small RESTful CRUD JSON API  microservice in Go for managing products in a store.

## Requirements

- Go 1.23
- docker
- docker-compose

## Implementation

This project is implemented using the following technologies:
- `postgres(v15)`  as the database
- `gin-gonic/gin` for the web framework
- `gorm.io/gorm` for the ORM
- `swaggo/swag` for the API documentation ({baseUrl}/swagger/index.html)
- `stretchr/testify` for the tests

### Good to know

The id of the products are UUIDs

### Available Routes

| Method | Endpoint                     | Data accepted                                                     | Description               |
|--------|------------------------------|-------------------------------------------------------------------|---------------------------|
| GET    | /api/v1/products             | query parameter **page** and **limit** (default: page=1 limit=10) | Get all products          |
| GET    | /api/v1/products/{id}        | path parameter **id**                                             | Get a product             |
| POST   | /api/v1/products/create      | json payload                                                      | Create product/s          |
| PUT    | /api/v1/products/update/{id} | path parameter **id**                                             | Update a product          |
| DELETE | /api/v1/products/delete/{id} | path parameter **id**                                             | Delete a product          |
| DELETE | /api/v1/products/delete      | json payload                                                      | Delete a list of products |

You can read more about the API in the API SWAGGER documentation  when you run the project locally  at `http://localhost:8080/swagger/index.html`

## HOW TO RUN

### Dockerized version

Make sure that in the .env file the `DB_HOST` is set with the value **pgsql**


1. Run the following command to build the go-app into a docker image:

```bash
docker build -t phrp/simple-crud:1.0 .
```

2. Then run the following command to run the docker-compose file that contains the go-app and the postgres database:

```bash
 docker-compose up
```

### Manually version

Make sure that in the .env file the `DB_HOST` is set with the value **localhost**

1. Run the following command to create a postgres database:

```bash
 docker-compose up pgsql  ## to run only the pgsql service
```
2. Run the following command to download and install the dependencies:

```bash
 go mod tidy
```
3. Run the following command to start the go-app:

```bash
 go run main.go
```

## TODO

- Add tests


