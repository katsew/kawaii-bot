.PHONY: init
init:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
	docker network rm kawaii-bot
	docker network create kawaii-bot

.PHONY: dev
dev:
	cd discord; source .envrc; docker-compose -f docker-compose.dev.yaml up -d
	cd heartcatch; source .envrc; docker-compose -f docker-compose.dev.yaml up -d

.PHONY: start
start:
	cd discord; source .envrc; docker-compose up -d
	cd heartcatch; source .envrc; docker-compose up -d

.PHONY: build
build:
	cd discord; source .envrc; docker-compose up -d --build
	cd heartcatch; source .envrc; docker-compose up -d --build

.PHONY: stop
stop:
	cd discord; docker-compose kill && docker-compose rm -f
	cd heartcatch; docker-compose kill && docker-compose rm -f

.PHONY: release
release: release_discord

.PHONY: release_discord
release_discord:
	docker build -t katsew/kawaii-bot/discord -f ./discord/Dockerfile .
	docker tag katsew/kawaii-bot/discord:latest 928962351838.dkr.ecr.ap-northeast-1.amazonaws.com/katsew/kawaii-bot/discord:latest
	docker push 928962351838.dkr.ecr.ap-northeast-1.amazonaws.com/katsew/kawaii-bot/discord:latest

.PHONY: release_hc
release_heartcatch:
	docker build -t katsew/kawaii-bot/heartcatch -f ./heartcatch/Dockerfile .
	docker tag katsew/kawaii-bot/heartcatch:latest 928962351838.dkr.ecr.ap-northeast-1.amazonaws.com/katsew/kawaii-bot/heartcatch:latest
	docker push 928962351838.dkr.ecr.ap-northeast-1.amazonaws.com/katsew/kawaii-bot/heartcatch:latest
