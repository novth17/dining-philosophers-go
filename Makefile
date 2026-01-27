BIN=bin/philo
PKG=./philo #go here to compile

build:
	go build -o $(BIN) $(PKG)

race:
	go build -race -o $(BIN) $(PKG)

run: build
	./$(BIN)

test:
	go test -v $(PKG) -count=1
