# Payments (payments-backend) - REST API in Go (Golang)

[![Build Status](https://travis-ci.org/michaljirman/payments-backend.svg?branch=master)](https://travis-ci.org/michaljirman/payments-backend)

[![GoDoc](https://godoc.org/github.com/michaljirman/payments-backend?status.svg)](https://godoc.org/github.com/michaljirman/payments-backend)

## REST API documentation available in folder ./src/paymentservice/docs

## Build native binaries
```
  cd ./src/paymentservice
  make
```

## Run tests
```
  cd ./src/paymentservice
  make test
```

## Run API 

```
  cd ./src/paymentservice
  ./bin/api
```

## CMD 

### Run CLI tool to list all payments

```
  cd ./src/paymentservice
  ./bin/findAll
```

## Docker

Build linux binaries:
```
  make linux-binaries
```

Build docker images and run API in a Docker container
```
  docker-compose up --build
```