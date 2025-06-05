run:
	go run .

test:
	go test -vet=off .

coverage:
	go test -vet=off -coverprofile=out/coverage.out .
	go tool cover -html=out/coverage.out

build:
	go build -o out/bin/trivia .
