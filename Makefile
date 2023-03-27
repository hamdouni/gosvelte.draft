ifndef VERBOSE
.SILENT:
endif
.PHONY: default restart start stop clean test build-client build-server build

default: start

restart: stop start

test:
	go test ./...

install:
	cd client && npm -s i

start:
	docker-compose up -d
	cd client && npm run dev

stop:
	echo "Stopping dev env"
	docker-compose stop 

clean:
	echo "cleaning application"
	rm -f server
	docker-compose down

build-client:
	echo "build client"
	cd client && npm -s run build

build-server:
	echo "build server"
	go build cmd/server.go

build: build-client build-server

