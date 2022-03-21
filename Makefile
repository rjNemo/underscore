TEST = "go test ./... -coverpkg=./... -coverprofile coverage.out -covermode=count; go tool cover -func coverage.out; rm coverage.out"

build:
	docker build -t underscore:latest .

test: build
	docker run --name underscore --rm -i -t underscore sh -c $(TEST)

doc:
	cd docs && hugo server -D