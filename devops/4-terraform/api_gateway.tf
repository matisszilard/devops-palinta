output "base_url" {
  value = aws_api_gateway_deployment.oath.invoke_url
}
