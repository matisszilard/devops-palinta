# Makefile with helper targets for deploying to AWS using Terraform

all:
	echo "Nothing here! Please choose other target!"

# Terraform commands

version?=latest
oath_app_version?=${version}
hero_app_version?=${version}
plan:
	terraform plan -var="app_version=${oath_app_version}" -var="hero_app_version=${hero_app_version}"

version?=latest
oath_app_version?=${version}
hero_app_version?=${version}
up: deploy-oath deploy-hero
	terraform apply -auto-approve -var="app_version=${oath_app_version}" -var="hero_app_version=${hero_app_version}"

version?=latest
hero_app_version?=latest
down:
	terraform destroy -var="app_version=${version}" -var="hero_app_version=${hero_app_version}"

# Prepare s3 buckets on aws

create-bucket:
	aws s3api create-bucket --bucket mszg-gondol --region=eu-central-1 --create-bucket-configuration LocationConstraint=eu-central-1

create-ui-bucket:
	aws s3api create-bucket --bucket mszg-gondol-ui --region=eu-central-1 --create-bucket-configuration LocationConstraint=eu-central-1

create-ui:
	aws s3 website s3://mszg-gondol-ui/ --index-document index.html --error-document error.html

# Build and deploy ui

url?=https://zj3ba3zqn2.execute-api.eu-central-1.amazonaws.com/dev
build-ui:
	rm -rf ./ui/index.html
	cat ./ui/index-template.html | sed 's#GONDOL_URL#${url}#' > ./ui/index.html

url?=https://zj3ba3zqn2.execute-api.eu-central-1.amazonaws.com/dev
sync-ui: build-ui
	aws s3 sync ui/ s3://mszg-gondol-ui/ --acl public-read

test-ui:
	open https://mszg-gondol-ui.s3.eu-central-1.amazonaws.com/index.html

# Build and deploy lambdas

build-lb: build-hero build-oath

build-oath:
	cd lambda/oath; GOOS=linux go build -o ../../build/main main.go; cd ../..
	zip -jrm build/oath.zip build/main

version?=latest
deploy-oath: build-oath
	aws s3 cp build/oath.zip s3://mszg-gondol/${version}/

build-hero:
	cd lambda/hero; GOOS=linux go build -o ../../build/main main.go; cd ../..
	zip -jrm build/hero.zip build/main

version?=latest
deploy-hero: build-hero
	aws s3 cp build/hero.zip s3://mszg-gondol/${version}/
