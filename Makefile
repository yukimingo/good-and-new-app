MIGRATION_PATH = ./backend/migrations
DB_URL = "mysql://user:pass@tcp(db:3306)/mydata?charset=utf8mb4&parseTime=True&loc=Local"

help:
	cat Makefile

up:
	docker compose up -d

upb:
	docker compose up --build -d

down:
	docker compose down

up-back:
	docker compose up backend -d

back:
	docker container exec -it good-and-new-backend-1 bash

db:
	docker container exec -it good-and-new-db-1 bash

mysql:
	docker compose exec db mysql -u user -p mydata

migrate:
	docker compose exec backend migrate -path ./migrations -database "mysql://user:pass@tcp(db:3306)/mydata?charset=utf8mb4&parseTime=True&loc=Local" up

go:
	docker compose exec backend go run cmd/main.go