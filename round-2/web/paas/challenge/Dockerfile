FROM golang:1.25 AS build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app cmd/main.go

FROM bash

RUN apk update && apk add iputils # For ping
COPY --from=build /go/bin/app /

USER 1001:1001

EXPOSE 8080
CMD ["/app"]
