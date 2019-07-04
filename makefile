.PHONY: doc

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tiyou .

run:
	go run main.go
