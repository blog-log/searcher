docker-build:
	docker build -t searcher:v0.0.1 .

docker-run:
	docker run -it -p 8080:8080 --env-file .env searcher:v0.0.1 

generate-secret-dev:
	kubectl create secret generic searcher-secret --dry-run=client --from-env-file=.env.dev -o yaml > k8s/secret-dev.yaml

generate-secret-prod:
	kubectl create secret generic searcher-secret --dry-run=client --from-env-file=.env.prod -o yaml > k8s/secret-prod.yaml