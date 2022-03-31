TEST = "go test ./... -coverpkg=./... -coverprofile coverage.out -covermode=count; go tool cover -func coverage.out; rm coverage.out"
IMAGE=underscore

build:
	docker build -t $(IMAGE):latest .

test: build
	docker run --name $(IMAGE) --rm -i -t $(IMAGE) sh -c $(TEST)

scan:
	trivy --cache-dir .trivycache/ image --exit-code 0  --no-progress --severity CRITICAL $(IMAGE)

scan-config:
	trivy config .

docs:
	cd docs && hugo server -D

build-docs:
	cd docs && hugo --gc --minify