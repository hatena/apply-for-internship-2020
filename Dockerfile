FROM golang:1.14-buster as builder

WORKDIR /go/src/github.com/hatena/apply-for-internship-2020/

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
RUN go build

FROM debian:buster-slim

WORKDIR /root/
COPY public.pem private.pem ./
COPY --from=builder /go/src/github.com/hatena/apply-for-internship-2020/apply-for-internship-2020 .

ENTRYPOINT ["./apply-for-internship-2020"]
