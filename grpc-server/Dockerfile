FROM golang:1.16.14-alpine

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV HOSTIP_KAFKA 192.168.1.10:9092

RUN go build -o /goServer

EXPOSE 50051

CMD ["/goServer"]