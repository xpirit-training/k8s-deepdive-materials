
install-mongo: 
	if ! kubectl get ns mongo; then kubectl create ns mongo; fi
	kubectl apply -f ./manifests/ebs-gp2-storageclass.yaml
	kubectl apply -f ./manifests/mongo-statefulset.yaml -n mongo

