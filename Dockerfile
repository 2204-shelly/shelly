FROM golang:alpine AS build

RUN mkdir /src /dist

WORKDIR /src
COPY . ./

ENV CGO_ENABLED=0
RUN go build -o /dist/shelly


# We're nice people, we give Ubuntu 20.04
FROM ubuntu:20.04
ARG DEBIAN_FRONTEND=noninteractive

# We're even nicer people, we give some tools
RUN apt-get update \
	&& apt-get install -y curl wget git g++ gcc libc6-dev make pkg-config iproute2 \
	&& rm -rf /var/lib/apt/lists/*

COPY --from=build /dist/shelly /usr/local/bin/shelly

ENTRYPOINT ["/usr/local/bin/shelly"]
