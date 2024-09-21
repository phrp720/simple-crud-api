# RESTful CRUD JSON API  microservice


### Description

A small RESTful CRUD JSON API  microservice in Go for managing products in a store.


### Routes


| Method | Endpoint              | Data accepted                                                     | Description                |
|--------|-----------------------|-------------------------------------------------------------------|----------------------------|
| GET    | /products             | query parameter **page** and **limit** (default: page=1 limit=10) | Get all products           |
| GET    | /products/{id}        | path parameter **id**                                             | Get a product              |
| POST   | /products/create      | json payload                                                      | Create  product/s          |
| PUT    | /products/update/{id} | path parameter **id**                                             | Update a product           |
| DELETE | /products/delete/{id} | path parameter **id**                                             | Delete a product           |
| DELETE | /products/delete      | json  payload                                                     | Delete a list of products  | 


