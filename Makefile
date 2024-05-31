# Makefile을 사용해 각 작업을 관리.
# make up 명령을 실행하면 백그라운드에서 자동 새로고침 갭라 환경 실행.
.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## 배포용 도커 이미지 빌드.
	docker build -t werellel/todoapp:${DOCKER_TAG} \
		--target deploy ./

build-local: ## 로컬 환경용 도커 이미지 빌드.
	docker compose build --no-cache

up: ## 자동 새로고침을 사용한 실행
	docker compose up -d

down: ## 종료
	docker compose down

logs: ## 로그 출력
	docker compose logs -f 

ps:
	docker compose ps

test:
	go test -race -shuffle=on ./...

help: ## 옵션 보기
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'