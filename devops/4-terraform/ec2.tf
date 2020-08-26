# Create ECP2

# Create a VPC to launch our instances into
resource "aws_vpc" "default" {
  cidr_block = "10.0.0.0/16"
}

# Create an internet gateway to give our subnet access to the outside world
resource "aws_internet_gateway" "default" {
  vpc_id = aws_vpc.default.id
}

# Grant the VPC internet access on its main route table
resource "aws_route" "internet_access" {
  route_table_id         = aws_vpc.default.main_route_table_id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.default.id
}

# Create a subnet to launch our instances into
resource "aws_subnet" "default" {
  vpc_id                  = aws_vpc.default.id
  cidr_block              = "10.0.1.0/24"
  map_public_ip_on_launch = true
}

# Our default security group to access
# the instances over SSH and HTTP
resource "aws_security_group" "default" {
  name        = "terraform_gondol"
  description = "Used in the terraform"
  vpc_id      = aws_vpc.default.id

  # SSH access from anywhere
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # HTTP access from the VPC
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/16"]
  }

  # outbound internet access
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_key_pair" "gondol" {
  key_name   = "gondolkey"
  public_key = file("./terraform.pub")
}

resource "aws_instance" "gondol" {
  key_name      = aws_key_pair.gondol.key_name
  ami           = "ami-0c115dbd34c69a004"
  instance_type = "t2.micro"
  vpc_security_group_ids = ["sg-07ff95d725c1c17a2"]
  subnet_id              = "subnet-03d07b29a9f559d89"
  iam_instance_profile = "fozocske-allow-s3"

 connection {
    type        = "ssh"
    user        = "ec2-user"
    private_key = file("./terraform")
    host        = self.public_ip
  }

  provisioner "remote-exec" {
    inline = [
      "sudo amazon-linux-extras enable nginx1.12",
      "sudo yum -y install nginx",
      "sudo aws s3 cp s3://mszg-gondol-ui /usr/share/nginx/html/ --recursive",
      "sudo systemctl start nginx"
    ]
  }
}
