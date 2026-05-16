CONFIG_FILE := configs/config.json

get_json = $(shell python3 -c "import json; f=open('$(CONFIG_FILE)'); data=json.load(f); print(data$(1))")

POSTGRES_USER := $(call get_json,['postgres']['user'])
POSTGRES_PASSWORD := $(call get_json,['postgres']['password'])
POSTGRES_PORT := $(call get_json,['postgres']['port'])
POSTGRES_DB := $(call get_json,['postgres']['dbname'])

export POSTGRES_USER POSTGRES_PASSWORD POSTGRES_DB

MIGRATE_DIR := database/migrations
DSN := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

.PHONY: info
info:
	@echo "Postgres User: $(POSTGRES_USER)"
	@echo "Postgres Password: $(POSTGRES_PASSWORD)"
	@echo "Postgres Port: $(POSTGRES_PORT)"
	@echo "Postgres DB: $(POSTGRES_DB)"
	@echo "DSN: $(DSN)"

.PHONY: install
install:
	@./scripts/install-tools.sh

.PHONY: rest
rest:
	go run cmd/rest/main.go
