ifndef VERBOSE
.SILENT:
endif
.PHONY: default restart start stop clean

default: start

restart: stop start

install:
	cd client && npm -s i

start:
	docker-compose up -d
	cd client && npm run dev &

stop:
	echo "Stopping dev env"
	docker-compose stop 

clean:
	echo "cleaning application"
	docker-compose down

