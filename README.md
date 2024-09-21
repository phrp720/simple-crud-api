# RESTful CRUD JSON API  microservice


### Description

A small RESTful CRUD JSON API  microservice in Go for managing products in a store.

### Routes

| Method | Endpoint                     | Data accepted                                                     | Description               |
|--------|------------------------------|-------------------------------------------------------------------|---------------------------|
| GET    | /api/v1/products             | query parameter **page** and **limit** (default: page=1 limit=10) | Get all products          |
| GET    | /api/v1/products/{id}        | path parameter **id**                                             | Get a product             |
| POST   | /api/v1/products/create      | json payload                                                      | Create product/s          |
| PUT    | /api/v1/products/update/{id} | path parameter **id**                                             | Update a product          |
| DELETE | /api/v1/products/delete/{id} | path parameter **id**                                             | Delete a product          |
| DELETE | /api/v1/products/delete      | json payload                                                      | Delete a list of products |

## TODO

- Add tests
- Add extra documentation and instructions
- Dockerization
