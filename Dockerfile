FROM golang:1.13

WORKDIR /go/src/docs

ENV GO111MODULE=on
ENV POSTGRESQL_IP_ADDRESS=172.20.1.3
ENV POSTGRESQL_PORT=5432
ENV POSTGRESQL_USER=postgres
ENV POSTGRESQL_PASSWORD=password
ENV POSTGRESQL_DBNAME=docs
ENV POSTGRESQL_MIGRATIONS_DIR=/go/src/docs/internal/migrations/

COPY cmd cmd
COPY configs configs
COPY internal internal
COPY Makefile Makefile
COPY go.mod go.mod
COPY go.sum go.sum

EXPOSE 80
ENTRYPOINT ["make"]