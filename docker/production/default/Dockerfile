FROM golang:1.13.0-alpine3.10 AS builder

ARG SERVICE_NAME=default

WORKDIR /go/src/github.com/abyssparanoia/catharsis-gcp/

ENV CGO_ENABLED=0

COPY . .

RUN go install -v -tags netgo -ldflags '-extldflags "-static"' ./cmd/default/


FROM alpine AS server

RUN apk add ca-certificates
COPY --from=builder /go/bin/default /bin/server
COPY --from=builder /go/src/github.com/abyssparanoia/catharsis-gcp/.env.default /go/src/github.com/abyssparanoia/catharsis-gcp/.env.default


WORKDIR /go/src/github.com/abyssparanoia/catharsis-gcp/

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ["/bin/server"]