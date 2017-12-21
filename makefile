
build: build/buddy_bot

build/buddy_bot: vendor $(shell find . -type f -name *.go)
	mkdir -pv build
	GOOS=linux GOARCH=amd64 go build -i -o build/buddy_bot main.go

vendor: Gopkg.lock
	dep ensure

docker/run:
	docker run buddy-bot:dev

docker/build/img:
	docker build -t buddy-bot:dev .
