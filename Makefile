build:
	make -j 2 build-fixtures build-server 
	
build-server:
	go build -x -o ./bin/out ./cmd/server

build-fixtures:
	go build -x -o ./bin/out ./cmd/fixture_db

gen:
	protoc --proto_path=api/proto --go_out=plugins=grpc:pkg/api api/proto/*.proto

clear:
	rm ./pkg/api/*.go

client: 
	evans -p 8080 api/proto/countries.proto

start:
	./bin/start.sh

stop:
	./bin/stop.sh

restart: 
	./bin/restart.sh
