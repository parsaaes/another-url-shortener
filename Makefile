export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export
export POSTGRES_ADDR=127.0.0.1:5432
export POSTGRES_USER=user
export POSTGRES_PASS=secret
export POSTGRES_DB=urlDB
export POSTGRES_DSN="postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@$(POSTGRES_ADDR)/$(POSTGRES_DB)?sslmode=disable"

migrate-install:
	which migrate || GO111MODULE=off go get -tags 'postgres' -v -u github.com/golang-migrate/migrate/cmd/migrate

migrate-create:
	migrate create -ext sql -dir migrations create_url_db

migrate-up:
	migrate -verbose -path ./migrations -database $(POSTGRES_DSN) up

migrate-down:
	migrate -path ./migrations -database $(POSTGRES_DSN) down

