run:
	go run ./cmd/docs

test:
	go test -v ./...

.DEFAULT_GOAL: run