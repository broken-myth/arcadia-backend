#!make
include .env

run:
	go run main.go

build:
	go build -ldflags "-w -s" -o arcadia_server main.go

start:
	./arcadia_server

watch:
	reflex -s -r '\.go$$' make run

lint:
	golangci-lint run

fix:
	golangci-lint run --fix

test:
	go run tests/run_tests.go

local_seed_dev:
	mysql -uroot -p${MYSQL_ROOT_PASSWORD} arcadia_23 < database/seed/dev_seeds.sql

local_seed_prod:
	mysql -uroot -p${MYSQL_ROOT_PASSWORD} arcadia_23 < database/seed/prod_seeds.sql

docker_seed_dev:
	bash scripts/restore_mysql_dump.sh database/seed/dev_seeds.sql

docker_seed_prod:
	bash scripts/restore_mysql_dump.sh database/seed/prod_seeds.sql

docker_test:
	docker exec -t arcadia_server go run tests/run_tests.go

docker_run:
	docker compose up --build -d

docker_down:
	docker compose down
