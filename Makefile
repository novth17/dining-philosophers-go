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

run-sched:
	go build -o bin/philo ./philo
		GODEBUG=schedtrace=1000 ./bin/philo 10000 800 200 200 > /dev/null
#GODEBUG=schedtrace=1000 ./bin/philo 5 800 200 200 > /dev/null

run-garbage:
	go build -o bin/philo ./philo
		GODEBUG=gctrace=1 ./bin/philo 10000 800 200 200 > /dev/null

test-short:
	go test -v -race -count=1 $(PKG) | grep -E "RUN|PASS|FAIL|died|time"

test:
	go test -v -count=1 $(PKG) | tee test_result.txt
#one stream goes to the file, and one stays in terminal

test-race:
	go test -v -race -count=1 $(PKG) | tee test_result.txt