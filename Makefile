build:
	go build -o bin/philo

run: build
	./bin/philo

test:
	go test -v ./... -count=1