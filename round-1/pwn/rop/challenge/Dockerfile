FROM ubuntu:22.04

RUN apt-get update
RUN apt-get install socat make gcc curl -y

EXPOSE 5000

RUN useradd ctf

WORKDIR /chal
COPY main.c /chal
COPY Makefile /chal
COPY flag.txt /chal

RUN make
RUN rm -rf Makefile
USER ctf

CMD ["socat", "tcp-l:5000,reuseaddr,fork", "EXEC:./pwn1"]