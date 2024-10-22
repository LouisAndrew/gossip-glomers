build:
	go build -o bin/server cmd/server/main.go

run-echo: build
	./bin/maelstrom/maelstrom test -w echo --bin ./bin/server --node-count 1 --time-limit 10

