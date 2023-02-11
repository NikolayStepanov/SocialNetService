up:
	docker-compose up
stop:
	docker-compose stop
down:
	docker-compose down
build:
	docker-compose build
swag:
	swag init -g cmd/app/main.go
build_app_db:
	docker compose -f app_db.yaml build app_db
run_app_db:
	docker compose -f app_db.yaml run -p 8080:8080 -d app_db
up_app_db:
	docker compose -f app_db.yaml up app_db
down_app_db:
	docker compose -f app_db.yaml down --remove-orphans
build_app_db_test:
	docker compose -f app_db_test.yaml build app_db_test
up_app_db_test:
	docker compose -f app_db_test.yaml up -d app_db_test
down_app_db_test:
	docker compose -f app_db_test.yaml down -v