FROM golang:1.12 AS builder
WORKDIR /go/src/app
COPY . .
RUN make server

FROM ubuntu:latest
RUN apt-get update
RUN apt-get install -y ca-certificates
WORKDIR /app/
COPY --from=builder /go/src/app/_output/bin/server .

ENTRYPOINT ["./server"]
CMD ["s"]