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

migrate:
	docker exec -it good-and-new-backend-1 sh -c "cd /app && go run cmd/migration/migrate.go"

go:
	docker exec -it good-and-new-backend-1 sh -c "cd /app && go run cmd/main.go"