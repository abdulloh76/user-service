.PHONY: clean build deploy

# To try different version of Go
GO := go

# Make sure to install aarch64 GCC compilers if you want to compile with GCC.
CC := aarch64-linux-gnu-gcc
GCCGO := aarch64-linux-gnu-gccgo-10

run-docker:
	docker-compose --env-file ./dev.env up
build-cmd:
	go build -o bin/main -v cmd/main.go
run-cmd:
	bin/main
