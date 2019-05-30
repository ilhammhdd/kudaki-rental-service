FROM golang:1.12-alpine AS build-env

RUN apk upgrade
RUN apk update
RUN apk add --no-cache curl
RUN apk add --no-cache git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/ilhammhdd/kudaki-rental-service
COPY . /go/src/github.com/ilhammhdd/kudaki-rental-service
RUN dep ensure
RUN go build -o kudaki_rental_service_app

FROM alpine 

ARG KAFKA_BROKERS
ARG REDISEARCH_SERVER
ARG DB_PATH
ARG DB_USERNAME
ARG DB_PASSWORD
ARG DB_NAME
ARG KAFKA_VERSION

ENV KAFKA_BROKERS=$KAFKA_BROKERS
ENV REDISEARCH_SERVER=$REDISEARCH_SERVER
ENV DB_PATH=$DB_PATH
ENV DB_USERNAME=$DB_USERNAME
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_NAME=$DB_NAME
ENV KAFKA_VERSION=$KAFKA_VERSION

COPY --from=build-env /go/src/github.com/ilhammhdd/kudaki-rental-service/kudaki_rental_service_app .

ENTRYPOINT ./kudaki_rental_service_app