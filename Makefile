.PHONY:
.SILENT:

BINARY=./bin/api

build:
	go build -o ${BINARY} cmd/api/main.go

run: build
	${BINARY}

clean:
	rm -f ${BINARY}

