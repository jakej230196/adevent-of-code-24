build:
	go mod tidy
	go get -u
	go build

run:
	go build
	./02

test:
	go test -v

lint:
	golangci-lint run --enable-all --disable gomnd --disable gochecknoglobals --disable nlreturn --fix ./...
