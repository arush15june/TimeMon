.PHONY: build

all: pre-build build

build: coordinator follower

pre-build: timeservice-proto

timeservice-proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/timeservice/timeservice.proto 

coordinator:
	go build -o ./build/coordinator ./cmd/coordinator/main.go

follower:
	go build -o ./build/follower ./cmd/follower/main.go
