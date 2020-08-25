output "base_url" {
  value = aws_api_gateway_deployment.oath.invoke_url
}

resource "null_resource" "deploy_ui" {
    provisioner "local-exec" {
        command = "make sync-ui url=${aws_api_gateway_deployment.oath.invoke_url}"
        }
  depends_on = [aws_api_gateway_deployment.oath]
}
