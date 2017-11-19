.PHONY: dev
dev:
	docker-compose -f docker-compose.dev.yaml up -d

.PHONY: run
run:
	docker-compose up -d

.PHONY: build
build:
	docker-compose build

.PHONY: stop
stop:
	docker-compose kill && docker-compose rm -f
