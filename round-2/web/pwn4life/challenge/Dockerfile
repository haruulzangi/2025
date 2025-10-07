FROM ubuntu:24.04 AS builder

WORKDIR /app

RUN apt-get update && apt-get install make gcc -y && rm -rf /var/lib/apt/lists/*

COPY . .
RUN make

FROM ubuntu:24.04

RUN apt-get update && apt-get install -y socat && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder --chmod=555 /app/pwn4life /app/challenge
COPY --chmod=111 entrypoint.sh /

EXPOSE 1337
CMD ["/entrypoint.sh"]
