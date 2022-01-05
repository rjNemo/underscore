TEST = "go test ./... -coverpkg=./... -coverprofile cp.out -covermode=atomic; go tool cover -func cp.out; rm cp.out"

build:
	docker build . -t underscore:latest

test: build
	docker run --name underscore --rm -i -t underscore bash -c $(TEST)

doc:
	cd docs && hugo server -D