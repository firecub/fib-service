# fib-service
A service for Fibonacci number functions

## Setup
- Open linux terminal
- Clone this repo using Git
- CD into `fib-service`

## Build
- Open linux terminal
- CD into `fib-service` directory created in Setup
- Build the Docker image using Docker (or podman, etc.)

   `docker build -t fib-service-app .`

## Run
- Open linux terminal
- CD into `fib-service` directory created in Setup
- Launch a container with this image

   `docker run -d -p 8080:8080 --name fib fib-service-app`

## Test
- Open linux terminal
- Test the app using `curl`

    `curl -v http://localhost:8080/fib/number/10`

- You should see HTTP status 200 (OK) witha response of "55"
