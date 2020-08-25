terraform {
  required_providers {
    # Commented out. It kills the syntax highligther in vscode.
    # aws = {
    #   source = "hashicorp/aws"
    # }
  }
}

# Oath lambda function
resource "aws_lambda_function" "oath" {
  function_name = "mszg-gondol-oath"

  # The bucket name as created earlier with "aws s3api create-bucket"
  s3_bucket = "mszg-gondol"
  s3_key    = "${var.app_version}/oath.zip"

  handler     = "main"
  runtime     = "go1.x"
  memory_size = 128

  role = aws_iam_role.lambda_exec.arn
}

# Hero lambda function
resource "aws_lambda_function" "hero" {
  function_name = "mszg-gondol-hero"

  s3_bucket = "mszg-gondol"
  s3_key    = "${var.hero_app_version}/hero.zip"

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

resource "aws_iam_role_policy" "frontend_lambda_role_policy" {
  name   = "frontend-lambda-role-policy"
  role   = aws_iam_role.lambda_exec.id
  policy = data.aws_iam_policy_document.lambda_log_and_invoke_policy.json
}

data "aws_iam_policy_document" "lambda_log_and_invoke_policy" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]
    resources = ["*"]
  }

  statement {
    effect = "Allow"
    actions = ["lambda:InvokeFunction"]
    resources = ["arn:aws:lambda:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:function:*"]
  }
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
# Create AWS API Gateway
resource "aws_api_gateway_resource" "oath" {
  rest_api_id = aws_api_gateway_rest_api.oath.id
  parent_id   = aws_api_gateway_rest_api.oath.root_resource_id
  path_part   = "oath"
}

# Internet -----> API Gateway
resource "aws_api_gateway_method" "get_oath" {
  rest_api_id   = aws_api_gateway_rest_api.oath.id
  resource_id   = aws_api_gateway_resource.oath.id
  http_method   = "GET"
  authorization = "NONE"
}

# API Gateway ------> Lambda
# For Lambda the method is always POST and the type is always AWS_PROXY.
#
resource "aws_api_gateway_integration" "lambda" {
  rest_api_id = aws_api_gateway_rest_api.oath.id
  resource_id = aws_api_gateway_method.get_oath.resource_id
  http_method = aws_api_gateway_method.get_oath.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.oath.invoke_arn
}

# Enable cors using terraform module

// TODO test if it is required
module "apigateway-cors" {
  source  = "mewa/apigateway-cors/aws"
  version = "2.0.0"

  api = aws_api_gateway_rest_api.oath.id
  resource = aws_api_gateway_resource.oath.id
  methods = ["GET"]
}

# API gateway deployment

resource "aws_api_gateway_deployment" "oath" {
  depends_on = [
    aws_api_gateway_integration.lambda,
  ]

  rest_api_id = aws_api_gateway_rest_api.oath.id
  stage_name  = "dev"
}
