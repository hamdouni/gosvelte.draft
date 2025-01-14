ifndef VERBOSE
.SILENT:
endif
.PHONY: default restart start stop clean test build-client build-server build

default: start

restart: stop start

test:
	go test ./...

install:
	go mod tidy
	cd internal/web && npm -s i

start:
	tmux send-keys 'goconvey' C-m \; \
		split-window -h -l 66 \; send-keys 'watcher -run internal/cmd/server/*.go' C-m \; \
		split-window -h \; send-keys 'cd internal/web && npm run dev' C-m \;

stop:
	echo "Stopping dev env"
	docker-compose stop 

clean:
	echo "cleaning application"
	rm -f server
	docker-compose down

build-client:
	echo "build client"
	cd internal/web && npm -s run build

build-server:
	echo "build server"
	go build -o ./build/ internal/cmd/server/ 

build: build-client build-server

