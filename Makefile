.PHONY: deploy

export VERSION ?= $(shell git show -q --format=%h)

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o arumba app/api/main.go
	docker build -t bickyeric/arumba:$(VERSION) -f Dockerfile .

test:
	go test ./... -cover -count=1

push:
	docker push bickyeric/arumba:$(VERSION)

deploy:
	envsubst < deploy/template.yml > deploy.yml
	docker stack deploy --compose-file deploy.yml arumba
