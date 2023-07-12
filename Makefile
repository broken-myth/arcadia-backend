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

seed:
	go run database/seedDatabase/seed.go

test:
	go run tests/run_tests.go
