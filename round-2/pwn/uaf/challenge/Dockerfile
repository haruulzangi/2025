#FROM ubuntu:16.04 as build
FROM ubuntu:20.04 AS build

ARG name=uaf

RUN apt-get update -y; apt-get install build-essential -y
ADD $name.c /tmp/$name.c
ADD Makefile /tmp/Makefile
RUN cd /tmp/; make all

FROM ubuntu:20.04

ARG name=uaf
ARG port=30013

RUN apt-get update -y; apt-get install socat -y
COPY --from=build /tmp/$name /pwn/$name
COPY flag.txt /pwn/flag.txt
WORKDIR /pwn

EXPOSE $port
#RUN echo "exec socat -s TCP-LISTEN:$port,reuseaddr,fork EXEC:/pwn/$name,stderr" > /pwn/docker_entrypoint.sh
RUN echo "socat tcp-l:$port,fork exec:/pwn/$name,reuseaddr" > /pwn/docker_entrypoint.sh

ENTRYPOINT ["sh", "/pwn/docker_entrypoint.sh"]
