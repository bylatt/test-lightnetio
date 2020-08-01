.DEFAULT: test
.SILENT: test start stop

test:
	go test ./... -count=1

start:
	docker-compose up -d

stop:
	docker-compose down