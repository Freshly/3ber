image := "gcr.io/freshly-docker/3ber"
tag := "main"

deps:
	go mod vendor
	go mod tidy

fmt:
	gofmt -w .

build: deps fmt
	CROSS=true ./scripts/build.sh

test:
	go test

docker-build: deps fmt
	./scripts/docker-build.sh

docker-push:
	./scripts/docker-push.sh
