provider "aws" {
  region = "ap-south-1"
}

resource "aws_instance" "example_instance" {
  ami           = "ami-03f4878755434977f"
  instance_type = "t2.micro"

  key_name        = "Aws-Key"
  subnet_id       = "subnet-0016ec7362742ba4d"
  vpc_security_group_ids = ["sg-06c6d6da0a14590d5"]

  tags = {
    Name = "MyInstance"
  }
}

output "instance_id" {
  value = aws_instance.example_instance.id
}

