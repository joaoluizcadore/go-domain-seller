run: 
	go run cmd/domain-seller/main.go

build:
	go build -o bin/domain-seller cmd/domain-seller/main.go

test:
	go test -v ./...

clean:
	rm -rf bin