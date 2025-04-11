run: 
	go run main.go

test:
	go run gotest.tools/gotestsum@latest

test-verbose:
	go run gotest.tools/gotestsum@latest --format short-verbose

test-clean:
	go clean -testcache

test-cover:
	go run gotest.tools/gotestsum@latest -- -coverprofile=cover.out ./... 

test-cover-verbose:
	go run gotest.tools/gotestsum@latest --format short-verbose -- -coverprofile=cover.out ./...j

tidy: 
	go mod tidy
