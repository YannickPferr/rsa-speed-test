FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /rsa-speed-test

EXPOSE 8080

ENTRYPOINT [ "/rsa-speed-test" ]
