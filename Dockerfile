FROM golang:alpine

WORKDIR /myapp
COPY . .

RUN go build -o ./bin/main-app ./cmd/api

CMD ["/myapp/bin/main-app"]
EXPOSE 8080
