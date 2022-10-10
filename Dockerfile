FROM golang:1.19.2 AS builder

WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY ./pkg ./pkg
COPY ./main.go  ./

ENV CGO_ENABLED=0
RUN go build -o /build/main main.go

FROM scratch as tickets
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder  /build/main /bin/cmd
USER 1000
ENTRYPOINT ["/bin/cmd"]
