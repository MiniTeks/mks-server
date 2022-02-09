FROM golang:1.17

WORKDIR /build
ADD . /build/

RUN mkdir /tmp/cache
RUN CGO_ENABLED=0 GOCACHE=/tmp/cache go build -mod=vendor -v -o /tmp/mks-server .

FROM alpine:3.15

WORKDIR /app
COPY --from=0 /tmp/mks-server /app/mks-server
COPY ./config /app/config
COPY ./entrypoint.sh /app/entrypoint.sh
RUN chmod a+x /app/entrypoint.sh

ENTRYPOINT ["sh", "/app/entrypoint.sh" ]