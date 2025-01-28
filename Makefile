run: 
	go run main.go

test: 
	go test -cover ./...

tidy: 
	go mod tidy
