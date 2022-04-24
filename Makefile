test:
	docker-compose up -d
	sleep 60
	docker exec -it api_gotrue_1 gotrue migrate
	go test ./...
	docker-compose down

start:
	docker-compose up -d
	sleep 60
	docker exec -it api_gotrue_1 gotrue migrate

run:
	MYSQL_HOST="127.0.0.1" \
	MYSQL_USER="mysql" \
	MYSQL_PASS="mysql" \
	go run main.go

stop:
	docker-compose down
