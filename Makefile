build:
	go build ./cmd/example_api

run:
	make -s build
	./example_api
