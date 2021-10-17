
test-healthz:
	@curl -XGET -I http://localhost:8000/healthz

test-response:
	@curl -XGET -I  http://localhost:8000 -H "MyHeader: test"

run:
	@export VERSION="1.1.1" && go run main.go