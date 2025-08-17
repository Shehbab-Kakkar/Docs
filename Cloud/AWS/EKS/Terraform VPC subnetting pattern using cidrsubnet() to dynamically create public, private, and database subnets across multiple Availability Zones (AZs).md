smart Terraform VPC subnetting pattern using cidrsubnet() to dynamically create public, private, and database subnets across multiple Availability Zones (AZs) — great for scalable MLOps deployments.

Let’s break this down step by step with:

✅ What you're doing

📘 How cidrsubnet works

📄 Terraform code example

🧠 Detailed explanation of the logic

🚀 Why this is scalable and useful

✅ 1. What You're Doing

You have a vpc_cidr block (e.g., 10.0.0.0/16), and you're using cidrsubnet() with count.index to split it into:

Public subnets

Private subnets (shifted by + var.az_count)

DB subnets (shifted by + 2 * var.az_count)

This creates non-overlapping subnets for each AZ, across three tiers, without manually defining IPs.

📘 2. How cidrsubnet() Works
cidrsubnet(base_cidr_block, newbits, netnum)


base_cidr_block: your base (e.g., 10.0.0.0/16)

newbits: how many bits to add to subnet (e.g., 8 → /24)

netnum: the subnet index (e.g., 0, 1, 2, etc.)

Example:

cidrsubnet("10.0.0.0/16", 8, 0) → 10.0.0.0/24
cidrsubnet("10.0.0.0/16", 8, 1) → 10.0.1.0/24


Each subnet gets its own /24 range.

📄 3. Terraform Example (with your logic)
🔧 Variables (in variables.tf)
variable "vpc_cidr" {
  default = "10.0.0.0/16"
}

variable "az_count" {
  default = 3
}

📁 Subnets (in main.tf)
# Example: Public Subnets
resource "aws_subnet" "public" {
  count                   = var.az_count
  vpc_id                 = aws_vpc.main.id
  cidr_block             = cidrsubnet(var.vpc_cidr, 8, count.index)
  availability_zone      = data.aws_availability_zones.available.names[count.index]
  map_public_ip_on_launch = true
  tags = {
    Name = "public-${count.index}"
  }
}

# Private Subnets (shifted by az_count)
resource "aws_subnet" "private" {
  count                   = var.az_count
  vpc_id                  = aws_vpc.main.id
  cidr_block              = cidrsubnet(var.vpc_cidr, 8, count.index + var.az_count)
  availability_zone       = data.aws_availability_zones.available.names[count.index]
  tags = {
    Name = "private-${count.index}"
  }
}

# Database Subnets (shifted by 2 * az_count)
resource "aws_subnet" "database" {
  count                   = var.az_count
  vpc_id                  = aws_vpc.main.id
  cidr_block              = cidrsubnet(var.vpc_cidr, 8, count.index + 2 * var.az_count)
  availability_zone       = data.aws_availability_zones.available.names[count.index]
  tags = {
    Name = "db-${count.index}"
  }
}

🗃️ Data Block for AZs
data "aws_availability_zones" "available" {}

🧠 4. Detailed Explanation
Subnet Type	Index Formula	CIDR Output (Example)
Public	count.index	10.0.0.0/24, 10.0.1.0/24, 10.0.2.0/24
Private	count.index + var.az_count	10.0.3.0/24, 10.0.4.0/24, 10.0.5.0/24
Database	count.index + 2 * var.az_count	10.0.6.0/24, 10.0.7.0/24, 10.0.8.0/24

This ensures non-overlapping subnet blocks per tier, automatically.

You're effectively allocating subnet ranges per AZ, per tier:

AZ0 → public-0, private-0, db-0

AZ1 → public-1, private-1, db-1

AZ2 → public-2, private-2, db-2

🚀 5. Why This Rocks (Best Practices)

✅ Scalable: Add or reduce AZs just by changing var.az_count

✅ No IP collisions: Because each subnet tier is offset

✅ Maintainable: No hardcoded IPs, easy to read and extend

✅ Multi-AZ Ready: Supports HA deployments (great for MLOps)

✅ Automation Friendly: Can be used with for_each, modules, or Terraform workspaces

🔄 Optional: Output Subnet Lists
output "public_subnet_ids" {
  value = aws_subnet.public[*].id
}

output "private_subnet_ids" {
  value = aws_subnet.private[*].id
}

🧪 Want Bonus: Turn This into a Module?

Let me know if you want this wrapped in a Terraform module with variables like:

VPC CIDR

Subnet mask

Tiers (e.g., public/private/db)

Region and AZ filtering
