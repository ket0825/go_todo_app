.PHONY: help build build-local up down logs ps test generate migrate dry-migrate
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## 배포용 도커 이미지 빌드 
		docker build -t week5:${DOCKER_TAG} --target deploy ./

build-local: ## 로컬 환경용 도커 이미지 빌드
		docker compose build --no-cache

up: ## 자동 새로고침을 사용한 docker compose 실행
		docker compose up -d

down: ## docker compose 종료
		docker compose down

logs:
		docker compose logs -f

ps:
		docker compose ps

test: 
		go test -race -shuffle=on ./...

dry-migrate: ## Try migration
	mysqldef -u todo -p todo -h 127.0.0.1 -P 33306 todo --dry-run < ./_tools/mysql/schema.sql

migrate: ## Execute migration
	mysqldef -u todo -p todo -h 127.0.0.1 -P 33306 todo < ./_tools/mysql/schema.sql

SHELL := /bin/bash
generate: ## Generate codes
	go generate ./...

help: ## 옵션 보기
		@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
				awk 'BEGIN {FS = ":.*?## "}; {PRINTF "\033[36m%-20s\033[0m %s\n", $$1, $$2}'