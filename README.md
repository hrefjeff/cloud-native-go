# Cloud Native Go

[Cloud Native app development with Go](https://learning-cloud-native-go.github.io/docs/hello-world-server/) - Simple bookshelf app with Go, PostgreSQL, and Docker. API specifications are as follows:

| Functionality | Resource | Method name | HTTP Method | Route |
|---------------|----------|-------------|-------------|-------|
| API Health | health | Read | GET | /livez |
| List Books | book | List | GET | /v1/books |
| Create Book | book | Create | POST | /v1/books |
| Read Book | book | Read | GET | /v1/books/{id} |
| Update Book | book | Update | PUT | /v1/books/{id} |
| Delete Book | book | Delete | DELETE | /v1/books/{id} |

## Generate OAS 2.0

```sh
swag init -g cmd/api/main.go -o .swagger -ot yaml
```

## Add a book to the shelf and list all the books

```sh
# Add a book to the shelf
curl -X 'POST' \
  'http://localhost:8080/v1/books' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "author": "John Smith",
  "description": "A fascinating book about cloud native development",
  "image_url": "https://example.com/book-cover.jpg",
  "published_date": "2023-01-15",
  "title": "Cloud Native Go Programming"
}'

# Get at all the books on the shelf
curl -X 'GET' http://localhost:8080/v1/books
````

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

## Common troubleshooting commands

```sh
# Check if api is already running
lsof -i :8080
kill -9 <PID>

# View logs
docker-compose logs

# Access the database container
docker-compose exec db psql -U myapp_user -d myapp_db

# Run database migrations
docker-compose exec api /bin/migrate up
```
