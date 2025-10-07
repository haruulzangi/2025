FROM rust:latest AS builder

WORKDIR /app

COPY Cargo.toml Cargo.lock ./

COPY . .
RUN cargo build --release

FROM debian:trixie-slim

RUN apt-get update && apt-get install -y --no-install-recommends libcurl4-openssl-dev && rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/target/release/pwn4life-server /app

EXPOSE 8080
CMD ["/app"]
