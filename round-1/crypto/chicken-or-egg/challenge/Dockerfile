FROM golang:1.25 AS build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app cmd/main.go

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /go/bin/app /

ENV FLAG='HZ2025{s0_h@rd_onl1_g0l@ng_can_th1s_SIGNATURE}'
CMD ["/app"]
