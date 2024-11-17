run: 
	go run main.go

test: 
	go test ./internal/... ./cmd/...

tidy: 
	go mod tidy
