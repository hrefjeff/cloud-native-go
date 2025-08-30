# Cloud Native Go

[Cloud Native app development with Go](https://learning-cloud-native-go.github.io/docs/hello-world-server/)

## Project Structure

```sh
.
├── cmd
│   └── api
│       └── main.go # Main application entry point
├── bin
│   └── main-app    # Compiled binary
├── Dockerfile
└── docker-compose.yml
```

## Build and Run

### Docker

```sh
# Build the Docker image
docker build -t myapp .

# Run the application
docker run -p 8080:8080 myapp
```

### Docker Compose

```sh
# Build the Docker image
docker-compose build

# Run the application
docker-compose up
```

## Database migrations

```sh
go build -o ./bin/migrate ./cmd/migrate
go build -o ./bin/api ./cmd/api
./bin/migrate create create_books_table sql
./bin/migrate up
```
