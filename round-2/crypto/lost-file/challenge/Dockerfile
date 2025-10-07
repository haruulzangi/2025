FROM golang:1.25 AS build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app cmd/main.go

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build --chmod=111 /go/bin/app /

EXPOSE 2222
CMD ["/app"]
