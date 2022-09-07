# classified-ads

This is a simple classified-ads microservice in Go.

## Goal
1. A way to insert a new ad. An ad consists of a subject, a body, an optional price, and
an email address.
2. A way of fetching existing ads. It should be possible to sort the ads on the time they
were inserted and by their price.
3. A way of deleting a previously inserted ad.

## Running it

The service needs Mysql to run succsuccessfullyesfully. Let's use `docker-compose` to run those servers.

```bash
docker-compose up -d mysql
```

And now you can build and run the server:

```bash
go run main.go
```

## Creating and image with docker

To run as a service we can create an image and use docker-compose to start the containers. First the image:

```bash
docker build . -t service:latest
```

Now we can run all the servers together using `docker-compose`.

```bash
docker-compose up
```

## Accessing the API
### Create Ad
POST http://localhost:8000/api/ads

{
    "email": "test@test.com", 
    "subject": "random-subject", 
    "price": 100.00,
    "body": "random-body"
}

### Get all Ads
http://localhost:8000/api/ads
#### Sorting and ordering
http://localhost:8000/api/ads?sortBy={created_at, price}&orderBy={asc, desc}

### Delete one ad
http://localhost:8000/api/ads/{id}

## Stopping the services

Again with `docker-compose`:

```bash
docker-compose down
```

## Run tests:
Firstly we need mysql to run the tests
```bash
docker-compose up mysql
```

Run all tests
```bash
go test ./...
```
