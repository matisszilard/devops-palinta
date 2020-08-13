#  ____            _   _           _
# |  _ \    __ _  | | (_)  _ __   | |_    __ _
# | |_) |  / _` | | | | | | '_ \  | __|  / _` |
# |  __/  | (_| | | | | | | | | | | |_  | (_| |
# |_|      \__,_| |_| |_| |_| |_|  \__|  \__,_|
#

.PHONY: all kube-up kube-down oc-up oc-down build clean docker-demeter docker-generator docker

all: palinta-up

#  __   __
# |  \ |_  \  /
# |__/ |__  \/
#

run-jenkins-agent:
	java -jar ${DEVOPS}/jenkins/test-local-agent/agent.jar \
	-jnlpUrl ${JENKINS_AGENT_URL} \
	-secret ${JENKINS_AGENT_SECRET} \
	-workDir "${DEVOPS}/jenkins/test-local-agent"

run-jenkins-golang-agent-in-docker:
	docker run -it --rm \
	--env JENKINS_SECRET=${JENKINS_AGENT_GOLANG_DOCKER_SECRET} \
	--env JENKINS_TUNNEL=${JENKINS_TUNNEL} \
	--env JENKINS_AGENT_NAME=${JENKINS_AGENT_GOLANG_DOCKER_NAME} \
	--env JENKINS_AGENT_WORKDIR=${JENKINS_AGENT_GOLANG_DOCKER_WORKDIR} \
	--env JENKINS_URL=${JENKINS_URL} \
	mszg/jenkins-slave:v0.5.0

run-jenkins-agent-docker-in-docker:
	docker run -it --rm \
	--env JENKINS_SECRET=${JENKINS_AGENT_GOLANG_DOCKER_SECRET} \
	--env JENKINS_TUNNEL=${JENKINS_TUNNEL} \
	--env JENKINS_AGENT_NAME=${JENKINS_AGENT_GOLANG_DOCKER_NAME} \
	--env JENKINS_AGENT_WORKDIR=${JENKINS_AGENT_GOLANG_DOCKER_WORKDIR} \
	--env JENKINS_URL=${JENKINS_URL} \
	-v /var/run/docker.sock:/var/run/docker.sock \
	jenkins-agent:latest

#           __   __  __        __ ___  __  __
# |_/ /  \ |__) |_  |__) |\ | |_   |  |_  (_
# | \ \__/ |__) |__ | \  | \| |__  |  |__ __)
#

kube-up:
	kubectl apply -f ${DEVOPS}/devops-palinta/devops

kube-down:
	kubectl delete -f ${DEVOPS}/devops-palinta/devops

palinta-up:
	oc apply -f ${DEVOPS}/devops-palinta/devops/palinta/palinta-user.yaml -n msz-palinta
	oc apply -f ${DEVOPS}/devops-palinta/devops/palinta/palinta-device.yaml -n msz-palinta

palinta-down:
	oc delete -f ${DEVOPS}/devops-palinta/devops/palinta/palinta-user.yaml -n msz-palinta
	oc delete -f ${DEVOPS}/devops-palinta/devops/palinta/palinta-device.yaml -n msz-palinta

up: palinta-up

down: palinta-down

gp:
	oc apply -f ${DEVOPS}/devops-palinta/devops/gp -n msz-palinta

gp-down:
	oc delete -f ${DEVOPS}/devops-palinta/devops/gp -n msz-palinta

elastic-pvc:
	oc apply -f ${DEVOPS}/devops-palinta/devops/elk/pvc.yaml -n elk-monitor

elastic: elastic-pvc
	oc apply -f ${DEVOPS}/devops-palinta/devops/elk/elasticsearch.yaml -n elk-monitor

elastic-down:
	oc delete -f ${DEVOPS}/devops-palinta/devops/elk/elasticsearch.yaml -n elk-monitor

logstash:
	oc apply -f ${DEVOPS}/devops-palinta/devops/elk/logstash.yaml -n elk-monitor

kibana:
	oc apply -f ${DEVOPS}/devops-palinta/devops/elk/kibana.yaml -n elk-monitor

kibana-down:
	oc delete -f ${DEVOPS}/devops-palinta/devops/elk/kibana.yaml -n elk-monitor

run-logstash:
	docker run --rm -it -v ${DEVOPS}/devops-palinta/devops/elk/cfg/logstash.yml:/usr/share/logstash/config/logstash.yml -v ${DEVOPS}/devops-palinta/devops/elk/cfg/pipeline:/usr/share/logstash/pipeline/ -p 5044:5044 -p 5000:5000 logstash:7.8.1

ek: elastic kibana

ek-down: elastic-down kibana-down

elk: elastic logstash kibana

jenkins-pvc:
	oc apply -f ${DEVOPS}/devops-palinta/devops/jenkins/pvc.yaml -n jenkins

jenkins-pvc-down:
	oc delete -f ${DEVOPS}/devops-palinta/devops/jenkins/pvc.yaml -n jenkins

jenkins: jenkins-pvc
	oc apply -f ${DEVOPS}/devops-palinta/devops/jenkins/jenkins.yaml -n jenkins

jenkins-down:
	oc delete -f ${DEVOPS}/devops-palinta/devops/jenkins/jenkins.yaml -n jenkins

#  __
# |__)     . |  _|
# |__) |_| | | (_|
#

build-demeter:
	cd cmd/demeter;   GOOS=linux   GOARCH=amd64 go build -o ../../build/linux-amd64/demeter; cd ../..
	cd cmd/demeter;   GOOS=darwin  GOARCH=amd64 go build -o ../../build/macos-amd64/demeter; cd ../..

build-data-generator:
	cd cmd/data-generator; GOOS=linux   GOARCH=amd64 go build -o ../../build/linux-amd64/data-generator; cd ../..
	cd cmd/data-generator; GOOS=darwin  GOARCH=amd64 go build -o ../../build/macos-amd64/data-generator; cd ../..

build-device:
	cd cmd/device;   GOOS=linux   GOARCH=amd64 go build -o ../../build/linux-amd64/device; cd ../..
	cd cmd/device;   GOOS=darwin  GOARCH=amd64 go build -o ../../build/macos-amd64/device; cd ../..

build-user:
	cd cmd/user;   GOOS=linux   GOARCH=amd64 go build -o ../../build/linux-amd64/user; cd ../..
	cd cmd/user;   GOOS=darwin  GOARCH=amd64 go build -o ../../build/macos-amd64/user; cd ../..

build: clean build-demeter build-data-generator build-device build-user

clean:
	rm -rf build

test:
	go test ./...

deps:
	go mod download
	go mod verify

#  __
# |__)     . |  _|    _|  _   _ |   _  _
# |__) |_| | | (_|   (_| (_) (_ |( (- |
#

docker-demeter: build-demeter
	docker build --build-arg target=demeter -t demeter -f ./Dockerfile .

docker-generator: build-data-generator
	docker build --build-arg target=data-generator -t data-generator -f ./Dockerfile .

docker-device: build-device
	docker build --build-arg target=device -t device -f ./Dockerfile .

docker-user: build-user
	docker build --build-arg target=user -t user -f ./Dockerfile .

docker-build: docker-demeter docker-generator docker-device docker-user

#  __
# |__)      _ |_
# |    |_| _) | )
#

tag ?= latest
docker-push-device: clean docker-device
	docker tag device mszg/palinta-device:${tag}
	docker push mszg/palinta-device:${tag}

tag ?= latest
docker-push-user: clean docker-user
	docker tag user mszg/palinta-user:${tag}
	docker push mszg/palinta-user:${tag}

tag ?= latest
docker-push: build docker-build
	docker tag demeter mszg/palinta-demeter:${tag}
	docker push mszg/palinta-demeter:${tag}
	docker tag data-generator mszg/palinta-generator:${tag}
	docker push mszg/palinta-generator:${tag}
	docker tag device mszg/palinta-device:${tag}
	docker push mszg/palinta-device:${tag}
	docker push mszg/palinta-generator:${tag}
	docker tag user mszg/palinta-user:${tag}
	docker push mszg/palinta-user:${tag}
