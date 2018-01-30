FROM golang:1.9

RUN mkdir /app
ADD hello /app/hello
WORKDIR /app
EXPOSE 8080

ENTRYPOINT /app/hello
