## Simple TODO API

Implementation of the API for a simple TODO application.

## Prequisites

1. Go (https://golang.org/doc/install)
2. PostgreSQL (https://www.postgresql.org/docs/9.3/tutorial-install.html)

## How to Run

### Run migrations

1. Install migrate (https://github.com/golang-migrate/migrate)
2. Run `migrate -database DB_CONNECTION_URL -path ./migrations up 4`. Db connection url example: (postgres://username:password@localhost:5432/database_name)

### Start the server

1. Copy `.env.example` to a new file `.env`
2. Update the enviroment variables in the `.env` file
3. Run `go mod download` to install dependencies
4. Run `go run main.go` to start the server