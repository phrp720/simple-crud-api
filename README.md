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
- `stretchr/testify` && `gin-gonic/gin` for the tests

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

You can read more about the API in the API SWAGGER documentation  when you run the project locally  at [API SWAGGER documentation](http://localhost:8080/swagger/index.html)

## HOW TO RUN

### Dockerized version

1. Make sure that in the .env file the `DB_HOST` is set with the value **pgsql**

You can manually change the variable in the .env file or run the following command:
```bash
 sed -i 's/^DB_HOST=.*/DB_HOST=pgsql/' .env
```

2. Run the following command to build the go-app into a docker image:

```bash
docker build -t phrp/simple-crud:1.0 .
```

3. Then run the following command to run the docker-compose file that contains the go-app and the postgres database:

```bash
 docker-compose up
```

### Manually version

1. Make sure that in the .env file the `DB_HOST` is set with the value **localhost**

You can manually change the variable in the .env file or run the following command:
```bash
 sed -i 's/^DB_HOST=.*/DB_HOST=localhost/' .env
```

2. Run the following command to create a postgres database:

```bash
 docker-compose up pgsql  ## to run only the pgsql service
```
3. Run the following command to download and install the dependencies:

```bash
 go mod tidy
```
4. Run the following command to start the go-app:

```bash
 go run main.go
```

### Running Tests

To run the tests, run the following command:

```bash
 go test ./tests
```
The tests are individual, and they are using an in-memory database (SQLite3). After each test, we flush the in-memory database so we can have individual tests. Each test verifies if the endpoints work correctly.



