FROM golang:1.19.2-alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY .. .
ENV GO111MODULE="on"
RUN go mod tidy
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/backend

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/backend --port 8080