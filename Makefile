COMPOSE_FILE = apps/docker-compose.yml
GO_MODULES = apps/smart_home apps/temperature-api

.PHONY: up
up:
	docker compose -f $(COMPOSE_FILE) up -d --build

.PHONY: down
down:
	docker compose -f $(COMPOSE_FILE) down -v

.PHONY: ps
ps:
	docker compose -f $(COMPOSE_FILE) ps

.PHONY: logs
logs:
	docker compose -f $(COMPOSE_FILE) logs -f

.PHONY: fmt
fmt:
	- for module in $(GO_MODULES); do (cd $$module && go fmt ./...); done

.PHONY: build
build:
	- for module in $(GO_MODULES); do (cd $$module && go build ./...); done

.PHONY: lint
lint:
	cd apps/temperature-api && golangci-lint run ./...
