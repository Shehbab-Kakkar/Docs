packer {
  required_plugins {
    amazon = {
      version = ">= 1.0.0"
      source  = "github.com/hashicorp/amazon"
    }
  }
}

variable "region" {
  type    = string
  default = "us-east-1"
}

source "amazon-ebs" "ubuntu" {
  region                  = var.region
  source_ami_filter {
    filters = {
      name                = "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"
      root-device-type    = "ebs"
      virtualization-type = "hvm"
    }
    owners      = ["xxxx"] # Canonical
    most_recent = true
  }
  instance_type           = "t3.micro"
  ssh_username            = "ubuntu"
  ami_name                = "custom-ubuntu-{{timestamp}}"
}

build {
  name    = "aws-custom-ubuntu-image"
  sources = [
    "source.amazon-ebs.ubuntu"
  ]

  provisioner "shell" {
    inline = [
      "sudo apt-get update && sudo apt-get upgrade -y",
      "sudo apt-get install -y nginx",
      "echo 'Hello from custom AMI' | sudo tee /var/www/html/index.html"
    ]
  }
}
