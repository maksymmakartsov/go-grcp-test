build:
	go build -v ./cmd/server/

gen:
	protoc --proto_path=api/proto --go_out=plugins=grpc:pkg/api api/proto/*.proto

clear:
	rm ./pkg/api/*.go