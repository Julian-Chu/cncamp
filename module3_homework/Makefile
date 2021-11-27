
test-healthz:
	@curl -XGET -I http://localhost:8000/healthz

test-response:
	@curl -XGET -I  http://localhost:8000 -H "MyHeader: test"

run:
	@export VERSION="1.1.1" && go run main.go

docker-build:
	@docker build -f ./Dockerfile -t julianchu/cncamp-go-server:dev ../module2_homework

docker-push:
	@docker tag julianchu/cncamp-go-server:dev julianchu/cncamp-go-server:latest
	@docker image push julianchu/cncamp-go-server:latest

docker-run:
	@docker run -p 8080:8000 julianchu/cncamp-go-server:dev
