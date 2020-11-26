FROM golang:1.14.9-alpine

ADD . /app

WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

RUN chmod +x ./simple-api

EXPOSE 7081

CMD /app/simple-api

