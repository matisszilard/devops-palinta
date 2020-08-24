
all:
	echo "Nothing here!"

up:
	terraform apply -var="app_version=1.0.0"

down:
	terraform destroy

create-bucket:
	aws s3api create-bucket --bucket mszg-gondol-ui --region=eu-central-1 --create-bucket-configuration LocationConstraint=eu-central-1

create-ui:
	aws s3 website s3://mszg-gondol-ui/ --index-document index.html --error-document error.html

sync-ui:
	aws s3 sync ui/ s3://mszg-gondol-ui/ --acl public-read

test-ui:
	open https://mszg-gondol-ui.s3.eu-central-1.amazonaws.com/index.html
