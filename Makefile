
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=. --go_out=. proto/skkm/skkm.proto
    

.PHONY: build
build: proto

	go build -o skkm-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t skkm-srv:latest
