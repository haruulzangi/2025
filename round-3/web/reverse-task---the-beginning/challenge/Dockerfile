FROM rust:latest AS builder

WORKDIR /app

COPY Cargo.toml Cargo.lock ./

COPY . .
RUN cargo build --release

# See: https://labs.iximiuz.com/tutorials/gcr-distroless-container-images
FROM gcr.io/distroless/cc-debian12:nonroot

COPY --from=builder /app/target/release/reverse-task-beginning /app

EXPOSE 8080
ENTRYPOINT ["/app"]
