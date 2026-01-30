BIN=bin/philo
PKG=./philo

build:
	go build -o $(BIN) $(PKG)

race:
	go build -race -o $(BIN) $(PKG)

run: build
	./$(BIN)

run-race: race
	./$(BIN)

run-sched: build
	GODEBUG=schedtrace=1000 ./$(BIN)

test:
	go test -v -race $(BIN)