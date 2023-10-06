

install-mongo: 
	if ! kubectl get ns mongo; then kubectl create ns mongo; fi
	kubectl apply -f ./manifests/mongo-basic.yaml -n mongo