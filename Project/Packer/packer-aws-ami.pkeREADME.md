# 🖼️ Packer AWS AMI Builder

This project shows how to use [Packer](https://www.packer.io/) to build a **custom AWS AMI** using HCL2 syntax.

You can use this as a base for building server images with pre-installed software (e.g., Nginx, Docker, your app, etc.).

---

## 🚀 Prerequisites

- [Packer](https://www.packer.io/downloads) v1.7+ installed  
- [AWS CLI](https://aws.amazon.com/cli/) configured (`aws configure`)
- IAM user/role with permissions for EC2, AMI, EBS, VPC, etc.

---

## 📁 Files

- `packer-aws-ami.pkr.hcl` – Packer template for AWS AMI build (see below)

---

## 📦 Example Template: `packer-aws-ami.pkr.hcl`

```hcl
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
    owners      = ["099720109477"] # Canonical
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
```

---

## 🛠️ Instructions

### 1. Save the Packer template

Create a file named `packer-aws-ami.pkr.hcl` in your project folder and paste the code above.

---

### 2. Initialize Packer

This will download the required plugins (e.g., for AWS):

```bash
packer init .
```

---

### 3. Validate the Template

Check your template for errors:

```bash
packer validate .
```

---

### 4. Build the AMI

Build your custom AMI:

```bash
packer build .
```

- The build process will:
  - Start a temporary EC2 instance from an official Ubuntu AMI
  - Install Nginx and customize `/var/www/html/index.html`
  - Create a new AMI from the modified instance
  - Output the new AMI ID

---

### 5. Change AWS Region or Customizations (Optional)

- To use a different region, edit `default = "us-east-1"` or run with:

  ```bash
  packer build -var="region=us-west-2" .
  ```

- To add more software, edit/add shell provisioner steps.

---

## 📝 Notes

- Your AWS credentials must have permissions to create EC2 instances, AMIs, and related resources.
- For more advanced uses (e.g., EBS volume size, VPC settings, tags), consult [Packer AWS docs](https://developer.hashicorp.com/packer/plugins/builders/amazon/ebs).
- Clean up old AMIs and snapshots to avoid unwanted AWS charges.

---

## 🧑‍💻 Example Result

After a successful build, you will see output like:

```
==> amazon-ebs.ubuntu: AMI: ami-0abcdef1234567890
```

You can now use this AMI ID to launch EC2 instances with your customizations.

---

## 📚 References

- [Packer by HashiCorp](https://www.packer.io/)
- [Packer AWS Builder Docs](https://developer.hashicorp.com/packer/plugins/builders/amazon/ebs)
