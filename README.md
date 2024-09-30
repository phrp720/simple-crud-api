# RESTful CRUD JSON API  microservice


## Description

A small RESTful CRUD JSON API  microservice in Go for managing products in a store.

## Requirements

- Go 1.23
- docker
- docker-compose

## Project Structure

### `simple-crud-api`
This folder contains the main application code, including handlers, repositories, and data access objects (DAOs).

- **`handler`**: Contains the HTTP handlers that define the API endpoints and their logic.
- **`router`**: Contains the API endpoints of the Application.
- **`repository`**: Contains the repository layer that interacts with the database.
- **`dao`**: Contains the data access objects that represent the database models.
- **`docs`**: Contains the API docs that are generated for swagger.
- **`db`**: Contains the database connection for the application.
- **`tests`**: Contains the test files for the application. It includes integration tests to ensure the functionality of the application.
  
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
| PUT    | /api/v1/products/update/{id} | path parameter **id** and json payload                            | Update a product          |
| DELETE | /api/v1/products/delete/{id} | path parameter **id**                                             | Delete a product          |
| DELETE | /api/v1/products/delete      | json payload                                                      | Delete a list of products |
| GET    | /swagger/index.html          | -                                                                 | API documentation         |

In the Swagger documentation, you will find all the endpoints along with the data they accept. Additionally, examples of the data that the endpoints accept and return are provided. You can also test the endpoints directly within the Swagger documentation. For more details, please refer to the API Swagger documentation available locally  at [API SWAGGER documentation](http://localhost:8080/swagger/index.html) .


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



