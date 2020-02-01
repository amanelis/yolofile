b: build

build:
	@go build -o bin/yolo main.go

# TESTING ----------------------------------------------------------------------
test: test_richgo

test_coverage_func:
	@go tool cover -func=coverage.out

test_coverage_html:
	@go tool cover -html=coverage.out

test_golang:
	@go test -v ./... -cover -coverprofile=coverage.out #-bench=.

test_gotest:
	@gotest -v ./... -cover -coverprofile=coverage.out #-bench=.

test_richgo:
	@richgo test -v ./... -cover -coverprofile=coverage.out #-bench=.
