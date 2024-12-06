build:
	go mod tidy
	go get -u
	go build

run:
	go build
	./01

test:
	go test -v

lint:
	golangci-lint run --enable-all --disable gomnd --disable gochecknoglobals --disable nlreturn --fix ./...
