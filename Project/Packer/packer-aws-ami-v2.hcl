packer {
  required_plugins {
    amazon = {
      version = ">= 1.2.8"
      source  = "github.com/hashicorp/amazon"
    }
  }
}

# Name of the package to install (e.g., nginx, apache2, docker.io)
variable "package" {
  description = "The package to install on the AMI"
}

# AWS region (e.g., us-west-2)
variable "region" {
  type        = string
  description = "AWS region to build the AMI in"
}

# EC2 Instance type for building (default: t3.micro)
variable "instancetype" {
  type    = string
  default = "t3.micro"
}

locals {
  timestamp = regex_replace(timestamp(), "[- TZ:]", "")
}

source "amazon-ebs" "ubuntu" {
  ami_name      = "demo-ubuntu-${var.package}-image-${local.timestamp}"
  instance_type = var.instancetype
  region        = var.region
  source_ami_filter {
    filters = {
      name                = "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"
      root-device-type    = "ebs"
      virtualization-type = "hvm"
    }
    most_recent = true
    owners      = ["123456789012"] # Example AWS account ID, replace as needed
  }
  ssh_username = "ubuntu"
}

build {
  name = "demo-packer-image"
  sources = [
    "source.amazon-ebs.ubuntu"
  ]
  provisioner "shell" {
    environment_vars = [
      "MESSAGE=Welcome to your demo AMI built with Packer!"
    ]
    inline = [
      "echo Installing selected package: ${var.package}",
      "sudo apt-get update",
      "sudo apt-get install -y ${var.package}",
      "echo $MESSAGE > index.html",
      "sudo mv index.html /var/www/html/ || true",
      "sudo systemctl start ${var.package} || true"
    ]
  }
}
