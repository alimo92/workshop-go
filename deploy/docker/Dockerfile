# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.22 AS build-stage

WORKDIR /workshop-go

COPY go.mod go.sum deploy/docker/startup.sh ./
COPY ./src ./src

RUN CGO_ENABLED=0 GOOS=linux go build -o workshop ./src/cmd/workshop

# Deploy the application binary into a lean image
FROM golang:1.22 AS build-release-stage

WORKDIR /workshop-go

COPY --from=build-stage /workshop-go/workshop /workshop-go/startup.sh ./
COPY .env  ./.env

EXPOSE 8080

ENTRYPOINT ["bash", "-c", "./startup.sh"]
