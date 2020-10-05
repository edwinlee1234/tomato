FROM golang:1.15.2-alpine3.12

ADD . /go/src/app
COPY ./config.example.yaml /go/src/app/config.yaml 

WORKDIR /go/src/app

ENV PORT 8080
EXPOSE 8080

RUN go build -mod=vendor -o main .
CMD ["/go/src/app/main"]