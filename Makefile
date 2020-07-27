#  __
# |__)  _  | .  _  |_  _
# |    (_| | | | ) |_ (_|
#

.PHONY: all kube-up kube-down oc-up oc-down build clean xbuild docker-demeter docker-generator docker

all: oc-up

# Kubernetes and Openshift predefined commands

kube-up:
	kubectl apply -f ${DEVOPS}/devops-palinta/kube

kube-down:
	kubectl delete -f ${DEVOPS}/devops-palinta/kube

oc-up:
	oc apply -f ${DEVOPS}/devops-palinta/kube -n msz-palinta

oc-down:
	oc delete -f ${DEVOPS}/devops-palinta/kube -n msz-palinta

up: oc-up

down: oc-down

build:
	mkdir build

clean:
	rm -rf build

# Build rules for Palinta projects

build-demeter:
	cd cmd/demeter;   GOOS=linux   GOARCH=amd64 go build -o ../../build/linux-amd64/demeter; cd ../..
	cd cmd/demeter;   GOOS=darwin  GOARCH=amd64 go build -o ../../build/macos-amd64/demeter; cd ../..

build-data-generator:
	cd cmd/data-generator; GOOS=linux   GOARCH=amd64 go build -o ../../build/linux-amd64/data-generator; cd ../..
	cd cmd/data-generator; GOOS=darwin  GOARCH=amd64 go build -o ../../build/macos-amd64/data-generator; cd ../..

build-device:
	cd cmd/device;   GOOS=linux   GOARCH=amd64 go build -o ../../build/linux-amd64/device; cd ../..
	cd cmd/device;   GOOS=darwin  GOARCH=amd64 go build -o ../../build/macos-amd64/device; cd ../..

xbuild: clean build-demeter docker-generator docker-device

# Build rules for building socker images

docker-demeter: build-demeter
	docker build --build-arg target=demeter -t demeter -f ./Dockerfile .

docker-generator: build-data-generator
	docker build --build-arg target=data-generator -t data-generator -f ./Dockerfile .

docker-device: build-device
	docker build --build-arg target=device -t device -f ./Dockerfile .

docker-build: docker-demeter docker-generator docker-device

# Push Palinta images

tag ?= latest
docker-push-device: clean docker-device
	docker tag device mszg/palinta-device:${tag}
	docker push mszg/palinta-device:${tag}

tag ?= latest
docker-push: xbuild docker-build
	docker tag demeter mszg/palinta-demeter:${tag}
	docker tag data-generator mszg/palinta-generator:${tag}
	docker tag device mszg/palinta-device:${tag}
	docker push mszg/palinta-demeter:${tag}
	docker push mszg/palinta-generator:${tag}
	docker push mszg/palinta-device:${tag}
