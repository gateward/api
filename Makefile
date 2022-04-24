test:
	docker-compose up -d
	sleep 30
	docker exec -it api_gotrue_1 gotrue migrate
	go test ./...
	docker-compose down

dev:
	docker-compose up -d
	sleep 30
	docker exec -it api_gotrue_1 gotrue migrate

stop:
	docker-compose down
