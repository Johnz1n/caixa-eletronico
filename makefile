export PROJECT = caixa-eletronico


build: 
	docker-compose build

log:
	docker-compose logs -f

up:
	docker-compose up -d

e2e-test:
	docker build -f dockerfile.e2e -t $(PROJECT)/e2e:1.0 --build-arg ENV_FILE=e2e/environment.json .
	docker run --rm --name caixa-eletronico-tests --network caixa-eletronico_shared-network $(PROJECT)/e2e:1.0 run --environment="/etc/environment.json"  /etc/newman/collection.json

deploy: build up log
