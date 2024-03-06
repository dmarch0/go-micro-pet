include .env

MIGRATOR_TAG = migrator:latest

build_migrator:
	docker build --build-arg="app=migrations" -t $(MIGRATOR_TAG) -f build/package/Dockerfile .

migrate_init:
	docker run --network deployments_base --env-file .env $(MIGRATOR_TAG) db init

migrate_create:
	docker run --network deployments_base --env-file .env $(MIGRATOR_TAG) db create_go