dev:
	docker-compose -f dev.docker-compose.yml up -d --build --scale worker=3

build:
	docker-compose -f docker-compose.yml build

up:
	docker-compose -f docker-compose.yml up -d --scale worker=5

down:
	docker-compose -f docker-compose.yml down
