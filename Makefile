.PHONY: default help

APP_NAME = personal-finance

info: help

default: help

help:
	@echo 'These are common ${APP_NAME} commands used in varios situations:'
	@echo
	@echo 'Usage:'
	@echo '   make run                                        Run the project.'
	@echo '   make install                                    Install all project dependencies.'
	@echo '   make migration-sql NAME=<option> EXT=<option>   Make migration sql.'
	@echo '   make migrate-up    EXT=<option>                 Migrate up tables.'
	@echo '   make migrate-down  EXT=<option>                 Migrate down tables.'

migration-sql:
	@echo "Create migration: ${NAME}.${EXT}"
	go run main.go make:migration ${NAME} ${EXT}

migrate-up:
	@echo "Migrate migration files"
	go run main.go migrate:up

migrate-down:
	@echo "Migrate down migration files"
	go run main.go migrate:down

run:
	@echo "Run the project"
	go run main.go

