FROM ubuntu:22.04

RUN apt-get update
RUN apt-get install socat make gcc curl -y

EXPOSE 5003

RUN useradd ctf

WORKDIR /chal
COPY pwn4 /chal
COPY flag.txt /chal

USER ctf

CMD ["socat", "tcp-l:5003,reuseaddr,fork", "EXEC:./pwn4"]
