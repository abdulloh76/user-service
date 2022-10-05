# user-service

protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative handlers/userGrpc/userGrpc.proto


evans userGrpc.proto -p 8080