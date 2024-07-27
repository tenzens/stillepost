FROM alpine:3.17

WORKDIR /home

COPY stillepost_daemon .

ENV coord_bind=0.0.0.0:2379
ENV peer_bind=0.0.0.0:2380

ENTRYPOINT ./stillepost_daemon member --coord-bind=${coord_bind} --peer-bind=${peer_bind}
