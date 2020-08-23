all: build run

.PHONY: run
run:
	cd app && ./server

.PHONY: clean
clean:
	rm -rf app

.PHONY: build-api
build-api:
	${MAKE} -C api build

.PHONY: build-client
build-client:
	${MAKE} -C client build

.PHONY: build
build: clean build-api build-client
	mkdir app
	mv client/build app/build
	mv api/bin/server app/
	cp build/Procfile app/
	cp build/Dockerfile app/

.PHONY: run-api
run-api:
	${MAKE} -C api run

.PHONY: run-client
run-client:
	${MAKE} -C client run

.PHONY: docker-build
docker-build: build
	docker build -t fullstack:0.0.1 ./app

.PHONY: docker-run
docker-run: docker-build
	docker run --publish 8080:8080 -e PORT=8081 fullstack:0.0.1

.PHONY: release-staging
release-staging: docker-build
	docker tag fullstack:0.0.1 registry.heroku.com/test-docker-staging-rch/web
	docker push registry.heroku.com/test-docker-staging-rch/web
	heroku container:release web -a test-docker-staging-rch
