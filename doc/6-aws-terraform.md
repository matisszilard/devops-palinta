# AWS and Terraform

## Step 0: To read & watch

- https://www.terraform.io/intro/index.html

## Step 1: Tutorial: Get Started - AWS

https://learn.hashicorp.com/collections/terraform/aws-get-started?track=getting-started#getting-started

## Step 2: Tutorial: Serverless Applications with AWS Lambda and API Gateway

https://learn.hashicorp.com/tutorials/terraform/lambda-api-gateway

## Step 3: Create a demo app on AWS using Terraform

Requirements:

- Have at least 1 public lambda function
- Have at least 1 private lambda function
- Have a simple web interface to call the public lambdas
- Host the web site on S3
- Simulate a deploy to AWS using terraform without manual user interactions

**Important notes**

- Don't forget to enable the CORS headers using terraform (Hint: use available modules)
- You can use existing roles for lambda creations (general lambda access, and call lambda function from other lambda functions)
- Deploy simulations can be done using Terraform, AWS CLI, Makefile, shell script or even Jenkins pipelines

## Step 3.1: 

Extend deployment with EC2, VPC, Subnet, Internet gateway creations...

Example: https://github.com/terraform-providers/terraform-provider-aws/tree/master/examples/two-tier

Host the website created in step 3 on EC2.

## Step 4: Have a :beer:, have a kitkat! :tada:
