
all:
	echo "Nothing here! Please choose other target!"

# Terraform commands

version?=v0.1.0
up:
	terraform apply -var="app_version=${version}"

version?=v0.1.0
down:
	terraform destroy -var="app_version=${version}"

# Prepare s3 buckets on aws

create-bucket:
	aws s3api create-bucket --bucket mszg-gondol --region=eu-central-1 --create-bucket-configuration LocationConstraint=eu-central-1

create-ui-bucket:
	aws s3api create-bucket --bucket mszg-gondol-ui --region=eu-central-1 --create-bucket-configuration LocationConstraint=eu-central-1

create-ui:
	aws s3 website s3://mszg-gondol-ui/ --index-document index.html --error-document error.html

sync-ui:
	aws s3 sync ui/ s3://mszg-gondol-ui/ --acl public-read

test-ui:
	open https://mszg-gondol-ui.s3.eu-central-1.amazonaws.com/index.html

# Build and deploy lambdas

version?=v0.1.0
deploy-oath:
	cd lambda/oath; GOOS=linux go build -o ../../build/main main.go; cd ../..
	zip -jrm build/oath.zip build/main
	aws s3 cp build/oath.zip s3://mszg-gondol/${version}/
