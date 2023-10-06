

build: 
	docker build -t docustore-backend:v0.0.1 .

re-run: stop run

stop:
	docker stop docustore-backend
	docker stop mongodb-local
	docker rm mongodb-local
	docker rm docustore-backend
	docker network rm local-network

run:
	docker network create local-network
	docker run -d --name docustore-backend --network local-network -p 8080:8080 -v "./config/local-config/config.properties:/workspace/config/config.properties" -v "./config/local-secret/secret.properties:/secret/database-secret/secret.properties" docustore-backend:v0.0.1 
	docker run --name mongodb-local --network local-network -d -p 27017:27017 --tmpfs /data/db:rw,noexec,nosuid,size=1G mongo:latest
 
build-push: build
	docker tag docustore-backend:v0.0.1 pathogen/docustore-backend:v0.0.1
	docker push pathogen/docustore-backend:v0.0.1

deploy: 
	if ! kubectl get ns docustore; then kubectl create ns docustore; fi
	kubectl apply -f ./manifests/. -n docustore


secret:
	kubectl create secret generic docustore-backend-secret --from-file=./config/remote-secret/secret.properties
