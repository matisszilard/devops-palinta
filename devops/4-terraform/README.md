# Deploy serverless application with AWS lambda functions and API Gateway using Terraform

A great tutorial to start with: https://learn.hashicorp.com/tutorials/terraform/lambda-api-gateway

The `Makefile` contains the targets to deploy to AWS.

## Deploy the latest version to AWS

```sh
make up
```

Delete the created resources from AWS:

```sh
make down
```

> Note: for further commands please check the `Makefile` as reference.

