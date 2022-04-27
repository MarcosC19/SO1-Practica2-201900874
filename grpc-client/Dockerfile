FROM golang:1.16.14-alpine
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV IP_SERVER 192.168.1.10:50051
RUN go build -o /goClient
EXPOSE 8080
CMD ["/goClient"]