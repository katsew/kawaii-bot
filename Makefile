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
