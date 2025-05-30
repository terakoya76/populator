lint:
	@if [ -z `which golangci-lint 2> /dev/null` ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sudo sh -s -- -b $(go env GOPATH)/bin v1.64.8; \
	fi
	@gofmt -s -w .
	@golangci-lint run --timeout 2m

test: lint
	go test -race -v ./...
