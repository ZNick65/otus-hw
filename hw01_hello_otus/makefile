lint:
	golangci-lint run .
	go test -v -count=1 -race -timeout=1m .
	./test.sh # При наличии

run:
	go run main.go