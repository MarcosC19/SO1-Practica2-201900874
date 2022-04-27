FROM golang:1.16-alpine
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download
COPY . . 
ENV RABBIT_HOST 34.71.238.30
ENV HOSTIP_MONGO 35.239.101.118
RUN go build -o goRabbitSubscriber
CMD ["/goRabbitSubscriber"]