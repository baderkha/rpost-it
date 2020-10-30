
start:
	docker-compose -f docker/local/docker-compose.yml up -d
	make build-local
	cd dist && ./main

shutodwn:
	docker-compose -f docker/local/docker-compose.yml down

build-local:prepare-build-dir
	cd client/rpost-it && yarn && yarn build
	mkdir -p dist/client/rpost-it/build
	cp -r client/rpost-it/build/* dist/client/rpost-it/build
	cp env.local.json dist/env.json
	go build -o dist/main src/main.go


build-deploy:prepare-build-dir
	cp env.json dist/env.json
	GOOS=linux go build -o dist/main src/main.go
	docker build -t shortner-url-ahmad-baderkhan -f docker/deploy/dockerfile .
	docker save -o shortner-api.zip shortner-url-ahmad-baderkhan:latest

deploy:build-deploy
	@clear
	@echo DEPLOYING LOCAL ZIP TO SERVER .....
	@scp -i ~/.ssh/eab_ssh shortner-api.zip root@ssh.shrter.xyz:/home/docker 
	@clear
	@echo START THE NEW IMAGE .... 
	@ssh -i ~/.ssh/eab_ssh root@ssh.shrter.xyz "cd /home/docker && docker load -i ./shortner-api.zip; docker kill shortner-url-ahmad-baderkhan; docker container rm shortner-url-ahmad-baderkhan; docker run -d -p 8080:8080 --name shortner-url-ahmad-baderkhan shortner-url-ahmad-baderkhan:latest; docker ps"
	
prepare-build-dir:
	mkdir -p dist/client
	cp README.md dist/README.md

build:prepare-build-dir
	cp env.json dist/env.json

	
