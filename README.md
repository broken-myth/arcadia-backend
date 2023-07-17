# [ARCADIA - Backend](https://github.com/delta/arcadia-backend)

This is the backend for the Arcadia game. It is written in GoLang and uses the Gin framework for HTTP serving.

## Prerequisites
#### Dev:
1. [Go](https://go.dev/doc/install)
1. [Reflex](https://github.com/cespare/reflex)
1. [Golangci-lint](https://golangci-lint.run/usage/install/)
1. [Redis](https://redis.io/download/)
1. [Swagger](https://github.com/swaggo/swag)

#### Docker:
1. [docker](https://docs.docker.com/engine/installation) 
(To contribute you will also need go, swag and golangci-lint)

## Setup 
proceed sequentially from here and run applicable commands


##### All Modes:

1. Clone the Repository.

1. `cp .env.example .env` and change values, if necessary.

1. `cp config.example.json config.json` and change values, if necessary.

##### Only for Dev:

1. Enable git hooks by running `git config core.hooksPath .githooks`.

1. Configure .vscode/settings.json to use golangci-lint for linting.
```json
{
   "go.lintTool":"golangci-lint",
   "go.lintFlags": [
   "--fast"
   ],
   "go.lintOnSave": "package",
   "go.formatTool": "goimports",
   "go.useLanguageServer": true,
   "[go]": {
      "editor.formatOnSave": true,
      "editor.codeActionsOnSave": {
         "source.organizeImports": true
      }
   },
   "go.docsTool": "gogetdoc"
}
```

1. Run `go mod download` to download all go dependencies.

1. Ensure you have a MySQL database running and create a database with the name you specified in `.env` and `config.json`.

##### For Dev in DOCKER:

1. Uncomment and comment the appropriate lines in `docker-compose.yaml`

1. Make changes to `.env` (set APP_ENV to DEV) and `config.json` (copy over from DOCKER config).

1. Follow all subsequent steps in the "Only for Docker" section below.

#### Only for Docker:
1. `docker compose up --build`

1. Access Adminer at http://localhost:8080/ (or at the ADMINER_EXTERNAL_PORT)

1. MySQL volumes are present in ./docker_volumes/mysql/ and logs at ./docker_volumes/logs/

1. Create and restore MySQL dumps with the scripts in ./scripts/

1. `docker compose down` to stop all containers

## Seeding the Database

1. To seed the database with dummy data, run `make seed` or `make docker_seed`

## Build and Run (for Dev only)

1. Run `make watch` command to start the backend in development mode. This will automatically restart the backend when you make changes to the code.

1. Run `make build` command to build the backend for production and `make start` to start the binary file (`arcadia_server`).


## Rules for Contributing:

1. All code should be formatted using `make lint` or `make fix` before committing.

1. Commit messages should be according to [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

1. Use the Pull Request template when creating a PR.


## Testing:

1. All Testing functions should be of the form `<feature>-test.go`

1. Include the test function(s) in `run_tests.go`

1. Run `make test` or `make docker_test` to run all tests.

Note: Testing in this repo has been made Docker-friendly, so not all best practices for tests may apply (such as filename = `*_test.go` or using the `testing` package or `go test -v ./...`)

## OpenAPI Documentation

1. Specify the URL in which you would like to have the docs in `config.SwaggerURL`

1. Write comments above Controllers to make a OpenAPI docs for that route.

1. Refer [swaggo/swag](https://github.com/swaggo/swag) documentation and learn how to write comments above controllers.

1. Run `swag fmt` to format the comments

1. Run `swag init` to build docs from the comments
