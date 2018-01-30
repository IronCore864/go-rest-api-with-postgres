FROM golang:1.9

COPY . /go/src/hello/
WORKDIR /go/src/hello

RUN go get ./
RUN go build

EXPOSE 8080
ENTRYPOINT /go/src/hello/hello
