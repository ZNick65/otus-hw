h help:
	@echo "lint			- run lint"
	@echo "test			- run all tests"
# @echo "run			- run the app"

#lint - run lint
lint:
	golangci-lint run --fix
	gofmt -w -s ./..
	go test -v -race -count=100 -timeout=1m .
	#./test.sh # При наличии
.PHONY: dev

# test - run all tests
test:
	go test --race ./...
.PHONY: test