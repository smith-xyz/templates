FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY *.go ./

RUN go build -o /web-service-template

EXPOSE 8080

CMD ["/web-service-template"]
