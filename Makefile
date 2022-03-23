TEST = "go test ./... -coverpkg=./... -coverprofile coverage.out -covermode=count; go tool cover -func coverage.out; rm coverage.out"

build:
	docker build -t underscore:latest .

test: build
	docker run --name underscore --rm -i -t underscore sh -c $(TEST)

docs:
	cd docs && hugo server -D

build-docs:
	cd docs && hugo --gc --minify