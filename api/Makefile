# Go parameters
BINARY_NAME=myapp

build:
	go build -o bin/$(BINARY_NAME) -v src/*.go

run:
	go run src/*.go

exec:
	go build -o bin/$(BINARY_NAME)-v src/*.go
	./bin/$(BINARY_NAME)

clean:
	rm bin/$(BINARY_NAME)*

compile:
	# 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o bin/$(BINARY_NAME)-freebsd-386 src/*.go
	# MacOS
	GOOS=darwin GOARCH=386 go build -o bin/$(BINARY_NAME)-darwin-386 src/*.go
	# Linux
	GOOS=linux GOARCH=386 go build -o bin/$(BINARY_NAME)-linux-386 src/*.go
	# Windows
	GOOS=windows GOARCH=386 go build -o bin/$(BINARY_NAME)-windows-386 src/*.go
    # 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o bin/$(BINARY_NAME)-freebsd-amd64 src/*.go
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_NAME)-darwin-amd64 src/*.go
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux-amd64 src/*.go
	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_NAME)-windows-amd64 src/*.go

