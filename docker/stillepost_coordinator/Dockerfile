FROM alpine:3.17

WORKDIR /home

COPY stillepost_daemon .

RUN apk add iputils

ENV client_bind=0.0.0.0:2379

ENTRYPOINT ./stillepost_daemon coordinator --client-bind=${client_bind}
