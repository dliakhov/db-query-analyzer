.PHONY: run start_docker_compose

start_docker_compose:
	docker compose up -d

run:
	go run main.go httpservice
