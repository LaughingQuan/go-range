clear:
	rm -f grpc_server grpc_cli xmirror-config.yaml

run:
	- docker stop GoGoatApp GoGoatOpenLdap GoGoatMysql 2>/dev/null || true
	- docker rm GoGoatApp GoGoatOpenLdap GoGoatMysql 2>/dev/null || true
	- docker network rm goat-network 2>/dev/null || true
	docker-compose up -d --build

gen_proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative rpc/goat.proto

grpc: clear
	./xmirror-go build -a -o ./grpc_server ./cmd/grpc/server/
	./xmirror-go build -a -o ./grpc_cli ./cmd/grpc/client/

all: clear run
