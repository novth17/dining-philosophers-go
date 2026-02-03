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

test-short:
	go test -v -race -count=1 $(PKG) | grep -E "RUN|PASS|FAIL|died|time"

test:
	go test -v -count=1 $(PKG) | tee test_result.txt
#one stream goes to the file, and one stays in terminal
