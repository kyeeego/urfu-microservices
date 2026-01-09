FROM golang:1.25.5-alpine3.23

RUN mkdir /app
ADD go.mod /app
ADD go.sum /app
WORKDIR /app
RUN go mod download

ADD . /app

RUN go build -o main .

CMD ["/app/main"]