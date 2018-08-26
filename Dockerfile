FROM golang:1.8

MAINTAINER Reja Nur Rochmat rezanurrochmat3@gmail.com

RUN mkdir /my-app

ADD . /my-app/ 

WORKDIR /my-app

RUN go get -u github.com/gorilla/mux

RUN go build -o mainApp .

CMD ["/my-app/mainApp"]