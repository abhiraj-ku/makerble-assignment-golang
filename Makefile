MIGRATION_DIR=internal/repository/db/migrations
DB_URL=postgres://postgres:helloDummy69@localhost:5432/healthcare?sslmode=disable
MIGRATE=migrate -path ${MIGRATION_DIR} -database "${DB_URL}"

# Run the go app
run:
   go run cmd/health-app/main.go

# build the go binary
build:
		go build -o bin/health-app cmd/health-app/main.go 

up:
	  docker compose up 

down:
    docker compose down

mig-up: 
  ${MIGRATE} up

mig-down:
  ${MIGRATE} down

create-mig:
   migrate create -ext sql -dir ${MIGRATION_DIR} -seq init_schema