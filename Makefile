# DEV

build-dev:
	docker build -t gochat -f container/images/Dockerfile . && docker build -t turn -f containers/images/Dockerfile.turn .

clean-dev:
	docker-compose -f container/composes/dc.dev.yml.down

run-dev:
	docker-compose -f container/composes/dc.dev.yml up

