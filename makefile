
build: build/buddy_bot

build/buddy_bot: vendor $(shell find . -type f -name *.go)
	mkdir -pv build
	go build -o build/buddy_bot main.go

vendor: Gopkg.lock
	dep ensure