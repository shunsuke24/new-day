FROM golang:1.22-alpine

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

WORKDIR /go/src/app/

ENV GO111MODULE=on

COPY ./src/go.mod ./src/go.sum ./

RUN go mod download

COPY ./src/app/ .

EXPOSE 8080

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
