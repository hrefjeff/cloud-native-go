FROM golang:alpine

WORKDIR /myapp
COPY . .

RUN go build -o main-app main.go

CMD ["/myapp/main-app"]
EXPOSE 8080
