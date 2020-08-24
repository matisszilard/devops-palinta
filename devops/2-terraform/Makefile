
all:
	echo "Nothing here!"

up:
	terraform apply -var="app_version=1.0.0"

down:
	terraform destroy

zip:
	cd example; zip ../example.zip main.js; cd ..

upload-src: zip
	aws s3 cp example.zip s3://mszg-bucket/v1.0.0/example.zip

test:
	aws lambda invoke --region=eu-central-1 --function-name=ServerlessExample output.txt
