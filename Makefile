TEST=go test ./...
COVER=-coverpkg=./... -coverprofile cov.out -covermode=count; go tool cover -func cov.out; rm cov.out
IMAGE=underscore

build:
	docker build -t $(IMAGE):latest .

test: build
	docker run --name $(IMAGE) --rm -i -t $(IMAGE) sh -c "$(TEST) $(COVER)"

scan:
	trivy --cache-dir .trivycache/ image --exit-code 0  --no-progress --severity CRITICAL $(IMAGE)

scan-config:
	trivy config .

.PHONY: docs
docs:
	cd docs && hugo server -D

build-docs:
	cd docs && hugo --gc --minify
