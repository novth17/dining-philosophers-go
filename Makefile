BIN=bin/philo
PKG=./philo

build:
	go build -o $(BIN) $(PKG)

# Added -race to the build command for a dedicated race binary 
race:
	go build -race -o $(BIN) $(PKG)

run: build
	./$(BIN)

# New: Run with Race Detector active (Checks for "Ghost Meals") 
run-race: race
	./$(BIN)

# New: Run with Scheduler Trace (1000ms intervals) to see GMP in action 
run-sched: build
	GODEBUG=schedtrace=1000 ./$(BIN)

# Updated: Added -race to tests to catch concurrency bugs during table-driven tests 
test:
	go test -v -race $(PKG) -count=1