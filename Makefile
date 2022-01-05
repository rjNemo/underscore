TEST = "go test ./... -coverpkg=./... -coverprofile coverage.txt -covermode=atomic; go tool cover -func coverage.txt"

build:
	docker build . -t underscore:latest

test: build
	docker run --name underscore --rm -i -t underscore bash -c $(TEST)

doc:
	cd docs && hugo server -D