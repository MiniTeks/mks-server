FROM golang:latest

WORKDIR /build
ADD . /build/

RUN mkdir /tmp/cache
RUN CGO_ENABLED=0 GOCACHE=/tmp/cache go build -mod=vendor -v -o /tmp/mks-server .

FROM scratch

WORKDIR /app
COPY --from=0 /tmp/mks-server /app/mks-server

CMD ["/app/mks-server" "-kubeconfig="]