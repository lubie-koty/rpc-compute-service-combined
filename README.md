# README
## Getting started
During application setup, make sure that you're in the root of the project directory and run the following commands:

```bash
$ cp .env.example .env # edit the .env file with your own settings
$ make tidy # pull dependencies
$ make run # start the server, alternatively use `make run/live` for live reloading
```

## Tests

To run tests, use the following command:

```bash
$ make test # `make test/cover` for outputting a coverage report
```

## Configuration settings

App settings are managed by using `.env` file, containing environment variables that are read into your application in the `run()` function in the `main.go` file.

## Make commands

The `Makefile` in the project root contains commands to easily run common admin tasks:
|     |     |
| --- | --- |
| `make tidy` | Format all code using `go fmt` and tidy the `go.mod` file. |
| `make audit` | Run `go vet`, `staticheck`, `govulncheck`, execute all tests and verify required modules. |
| `make test` | Run all tests. |
| `make test/cover` | Run all tests and outputs a coverage report in HTML format. |
| `make build` | Build a binary for the `cmd/api` application and store it in the `/tmp/bin` folder. |
| `make run` | Build and then run a binary for the `cmd/api` application. |
| `make run/live` | Build and then run a binary for the `cmd/api` application (uses live reloading). |
