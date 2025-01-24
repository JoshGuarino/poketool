run: 
	go run main.go

test: 
	go test -cover ./internal/... ./cmd/...

tidy: 
	go mod tidy
