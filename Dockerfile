FROM golang:alpine AS build

RUN mkdir /src /dist

WORKDIR /src
COPY . ./

ENV CGO_ENABLED=0
RUN go build -o /dist/shelly


# We're nice people, we give Ubuntu 20.04
FROM ubuntu:20.04

COPY --from=build /dist/shelly /usr/local/bin/shelly

ENTRYPOINT ["/usr/local/bin/shelly"]
