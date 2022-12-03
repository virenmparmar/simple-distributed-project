proto:
	rm -f web/auth/pb/*.go
	protoc --proto_path=web/auth/proto --go_out=web/auth/pb --go_opt=paths=source_relative \
    --go-grpc_out=web/auth/pb --go-grpc_opt=paths=source_relative \
    web/auth/proto/*.proto
evans:
	evans --host localhost --port 9000 -r repl

.PHONY: proto evans