terraform {
  required_providers {
    # aws = {
    #   source = "hashicorp/aws"
    # }
  }
}

provider "aws" {
  region = "eu-central-1"
}

resource "aws_lambda_function" "oath" {
  function_name = "mszg-gondol-oath"

  # The bucket name as created earlier with "aws s3api create-bucket"
  s3_bucket = "mszg-gondol"
  s3_key    = "${var.app_version}/oath.zip"

  # "main" is the filename within the zip file (main.js) and "handler"
  # is the name of the property under which the handler function was
  # exported in that file.
  handler     = "main"
  runtime     = "go1.x"
  memory_size = 128

  role = aws_iam_role.lambda_exec.arn
}

# IAM role which dictates what other AWS services the Lambda function
# may access.
resource "aws_iam_role" "lambda_exec" {
  name = "serverless_oath_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

}

# Configure API gateway

resource "aws_api_gateway_rest_api" "oath" {
  name        = "mszg-gondol-oath"
  description = "mszg gondol oath endpoint"
}

# Allow API gateway to invoke the hello Lambda function.
resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.oath.function_name
  principal     = "apigateway.amazonaws.com"

  # The "/*/*" portion grants access from any method on any resource
  # within the API Gateway REST API.
  source_arn = "${aws_api_gateway_rest_api.oath.execution_arn}/*/*"
}

# A Lambda function is not a usual public REST API. We need to use AWS API
# Gateway to map a Lambda function to an HTTP endpoint.
resource "aws_api_gateway_resource" "proxy" {
  rest_api_id = aws_api_gateway_rest_api.oath.id
  parent_id   = aws_api_gateway_rest_api.oath.root_resource_id
  path_part   = "oath"
}

# Internet -----> API Gateway
resource "aws_api_gateway_method" "proxy" {
  rest_api_id   = aws_api_gateway_rest_api.oath.id
  resource_id   = aws_api_gateway_resource.proxy.id
  http_method   = "GET"
  authorization = "NONE"
}

# API Gateway ------> Lambda
# For Lambda the method is always POST and the type is always AWS_PROXY.
#
resource "aws_api_gateway_integration" "lambda" {
  rest_api_id = aws_api_gateway_rest_api.oath.id
  resource_id = aws_api_gateway_method.proxy.resource_id
  http_method = aws_api_gateway_method.proxy.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.oath.invoke_arn
}

# resource "aws_api_gateway_method" "proxy_root" {
#   rest_api_id   = aws_api_gateway_rest_api.oath.id
#   resource_id   = aws_api_gateway_rest_api.oath.root_resource_id
#   http_method   = "ANY"
#   authorization = "NONE"
# }

# resource "aws_api_gateway_integration" "lambda_root" {
#   rest_api_id = aws_api_gateway_rest_api.oath.id
#   resource_id = aws_api_gateway_method.proxy_root.resource_id
#   http_method = aws_api_gateway_method.proxy_root.http_method

#   integration_http_method = "POST"
#   type                    = "AWS_PROXY"
#   uri                     = aws_lambda_function.oath.invoke_arn
# }

# API gateway deployment

resource "aws_api_gateway_deployment" "oath" {
  depends_on = [
    aws_api_gateway_integration.lambda,
    # aws_api_gateway_integration.lambda_root,
  ]

  rest_api_id = aws_api_gateway_rest_api.oath.id
  stage_name  = "dev"
}

# Enable cors using terraform module
module "cors" {
  source = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.1"

  api_id          = aws_api_gateway_rest_api.oath.id
  api_resource_id = aws_api_gateway_method.proxy.resource_id
}
