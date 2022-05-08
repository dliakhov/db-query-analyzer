.PHONY: run start_docker_compose

start_docker_compose:
	docker compose up -d

migrate_db:
	go run main.go migrate

run: migrate_db
	go run main.go httpservice
