.PHONY: run migrateup addmigration

start_db:
	docker compose up -d

migrate_db: start_db
	go run main.go migrate

run: migrate_db
	go run main.go httpservice
