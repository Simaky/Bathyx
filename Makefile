run:
	go run main.go

lint:
	golangci-lint run

lint-install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

lint-install-mac:
	brew install golangci-lint
