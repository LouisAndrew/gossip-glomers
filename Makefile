build:
	go build -o bin/server cmd/server/main.go

run-echo: build
	./bin/maelstrom/maelstrom test -w echo --bin ./bin/server --node-count 1 --time-limit 10

run-unique-ids: build
	./bin/maelstrom/maelstrom test -w unique-ids --bin ./bin/server --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition
