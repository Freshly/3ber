image := "gcr.io/freshly-docker/3ber"
tag := "main"

deps:
	go mod vendor
	go mod tidy

fmt:
	gofmt -w .

build: deps fmt
	./scripts/build.sh

test:
	go test

build_docker: deps fmt
	docker build . \
		-t {{image}}:{{tag}}

publish_docker:
	docker push {{image}}:{{tag}}
