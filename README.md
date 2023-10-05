# product-service

# Product Service

This service manages all functions related to products such as searching and cataloguing.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.16 or newer)

### Installing Dependencies

```bash
go mod tidy
```

### Docker Build

```sh
docker build -t product-service .
```

### Run the Docker Container 

```sh
docker run -d -p 5001:5001 --name product-service-container product-service
```

### Test the application 

```sh
curl http://localhost:5001/product/1
```