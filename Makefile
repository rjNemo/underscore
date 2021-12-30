test:
	docker build .

doc:
	cd docs && hugo server -D