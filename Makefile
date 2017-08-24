default: install

install:
	@dep ensure
	@go build -o grpc-health ./cmd/
	@cp ./grpc-health $(GOPATH)/bin/
	@rm grpc-health

uninstall:
	@rm $(GOPATH)/bin/grpc-health
