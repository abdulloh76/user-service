.PHONY: clean build deploy

FUNCTIONS := authorizer signin signup

# To try different version of Go
GO := go

# Make sure to install aarch64 GCC compilers if you want to compile with GCC.
CC := aarch64-linux-gnu-gcc
GCCGO := aarch64-linux-gnu-gccgo-10

run-docker:
	docker-compose --env-file ./dev.env up
build-server:
	go build -o bin/server/main -v cmd/server/main.go
run-server:
	bin/server/main
build-grpc:
	go build -o bin/grpcserver/main -v cmd/grpcserver/main.go
run-grpc:
	bin/grpcserver/main


build-lambdas:
	${MAKE} ${MAKEOPTS} $(foreach function,${FUNCTIONS}, build-${function})

clean:
	@rm $(foreach function,${FUNCTIONS}, functions/${function}/bootstrap)


invoke-authorizer:
	@sam local invoke --env-vars env-vars.json --event functions/authorizer/event.json AuthMiddlewareFunction
invoke-signin:
	@sam local invoke --env-vars env-vars.json --event functions/signin/event.json SignInFunction
invoke-signup:
	@sam local invoke --env-vars env-vars.json --event functions/signup/event.json SignUpFunction

