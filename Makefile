lint:
	golangci-lint run --fix

test:
	go test -json ./... -count=1 -cover | gotestfmt