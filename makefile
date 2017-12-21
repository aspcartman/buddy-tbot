
build/buddy_bot: vendor
	mkdir -pv build
	go build -o build/buddy_bot main.go

vendor: Gopkg.lock
	dep ensure