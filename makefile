APP_NAME=banners_service
DB_CONTAINER=postgres

PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
$(shell mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)

GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint

.PHONY: .install-linter
.install-linter:
	@if [ ! -f $(GOLANGCI_LINT) ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v1.46.2; \
	fi

.PHONY: lint
lint: .install-linter
	$(GOLANGCI_LINT) run ./... --config=.golangci.yml

.PHONY: lint-fast
lint-fast: .install-linter
	$(GOLANGCI_LINT) run ./... --fast --config=.golangci.yml



all: build


build:
	docker build -t $(APP_NAME) .


up:
	docker-compose up --build
	
init-db:
	docker exec -it postgres psql -U admin -d postgres -c "CREATE DATABASE test_db;"


start: up init-db


down:
	docker-compose down


restart: down start


clean:
	docker-compose down --rmi all --remove-orphans
	docker rmi $(APP_NAME)


logs:
	docker-compose logs -f


test:
	docker exec -it $(APP_NAME) go test ./internal/handler/


shell:
	docker exec -it $(APP_NAME) /bin/sh


help:
	@echo "make start - полный запуск с инициализацией баз данных"
	@echo "make build - сборка Docker образа $(APP_NAME)"
	@echo "make up - запуск всех сервисов через Docker Compose"
	@echo "make down - остановка всех сервисов"
	@echo "make restart - перезапуск всех сервисов"
	@echo "make clean - удаление всех созданных Docker образов и контейнеров"
	@echo "make logs - вывод логов приложения"
	@echo "make test - запуск тестов в контейнере"
	@echo "make shell - запуск оболочки в контейнере приложения"
	@echo "make help - вывод этой справки"