FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o http-proxy-server

EXPOSE 8080

CMD [ "/app/http-proxy-server" ]

