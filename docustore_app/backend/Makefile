

build: 
	docker build -t docustore-backend:v0.0.1 .


docker-run:
	docker run -d --name docustore-backend -p 8080:8080 docustore-backend:v0.0.1

build-push: build
	docker tag docustore-backend:v0.0.1 pathogen/docustore-backend:v0.0.1
	docker push pathogen/docustore-backend:v0.0.1

deploy: 
	if ! kubectl get ns docustore; then kubectl create ns docustore; fi
	kubectl apply -f ./manifests/. -n docustore


secret:
	kubectl create secret generic docustore-backend-secret --from-file=./config/remote-secret/secret.properties
