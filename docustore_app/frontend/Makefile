

build: 
	docker build -t docustore-frontend:v0.0.1 .

docker-run:
	docker run -d --name docustore-frontend -p 3000:3000 docustore-frontend:v0.0.1

build-push: build
	docker tag docustore-frontend:v0.0.1 pathogen/docustore-frontend:v0.0.1
	docker push pathogen/docustore-frontend:v0.0.1

deploy: 
	if ! kubectl get ns docustore; then kubectl create ns docustore; fi
	kubectl apply -f ./manifests/. -n docustore

cleanup: 
	docker stop docustore-frontend
	docker rm docustore-frontend

rebuild-and-run: cleanup build docker-run