test:
	go test ./pkg/... . -v

build:
	go build -o ./bin/server ./cmd/server/main.go

run:
	./bin/server

run-local:
	./bin/server -path ~/.kube/config