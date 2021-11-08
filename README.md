# Multipurpose Product API

## Features

1. RESTful web services with relevant HTTP Verbs
2. GRPC bi-directional streaming
3. Graceful Shutdown
4. Swagger documentation with ReDoc
5. Unit testing
6. JSON Validation
7. CORS Handling
8. Multi-Part Request Handling

## Services

### Product API [./product-api](./golang-yt-microservices)

RESTful Go based JSON API built using the Gorilla framework. The API allows CRUD based operations on a product list.

### Frontend website [./frontend](./golang-yt-microservices-frontend)

ReactJS website for presenting the Product API information.

### Currency [./currency](./currency)

gRPC service supporting simple Unaray and Bidirectional streaming methods.

### Product Images [./product-images](./product-images)

Go based image service supporting Gzipped content, multi-part forms and a RESTful
approach for uploading and downloading images.
