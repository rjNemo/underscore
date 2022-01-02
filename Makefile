TEST = "go test ./... -v -coverpkg=./... -coverprofile cp.out; go tool cover -func cp.out; rm cp.out"

build:
	docker build . -t underscore:latest

test: build
	docker run --name underscore --rm -i -t underscore bash -c $(TEST)

doc:
	cd docs && hugo server -D