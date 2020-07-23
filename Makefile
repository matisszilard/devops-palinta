.PHONY: all kube-up kube-down oc-up oc-down build clean xbuild docker-demeter docker-generator docker

all: oc-up

kube-up:
	kubectl apply -f ${DEVOPS}/devops-palinta/kube

kube-down:
	kubectl delete -f ${DEVOPS}/devops-palinta/kube

oc-up:
	oc apply -f ${DEVOPS}/devops-palinta/kube -n msz-palinta

oc-down:
	oc delete -f ${DEVOPS}/devops-palinta/kube -n msz-palinta

build:
	mkdir build

clean:
	rm -rf build

xbuild: clean build
	cd cmd/data-generator; GOOS=linux   GOARCH=amd64 go build -o ../../build/linux-amd64/data-generator; cd ../..
	cd cmd/data-generator; GOOS=darwin  GOARCH=amd64 go build -o ../../build/macos-amd64/data-generator; cd ../..
	cd cmd/demeter;   GOOS=linux   GOARCH=amd64 go build -o ../../build/linux-amd64/demeter; cd ../..
	cd cmd/demeter;   GOOS=darwin  GOARCH=amd64 go build -o ../../build/macos-amd64/demeter; cd ../..

docker-demeter:
	docker build -t demeter  -f ./Dockerfile.demeter .

docker-generator:
	docker build -t data-generator  -f ./Dockerfile.generator .

docker-build: docker-demeter docker-generator

tag ?= latest
docker-push: clean xbuild docker-build
	docker tag demeter mszg/demeter:${tag}
	docker tag data-generator mszg/palinta-generator:${tag}
	docker push mszg/demeter:${tag}
	docker push mszg/palinta-generator:${tag}
