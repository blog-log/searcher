docker-build:
	docker build -t searcher:v0.0.1 .

docker-run:
	docker run -it -p 8080:8080 --env-file .env searcher:v0.0.1 
