# 🚀 Smart Terraform VPC Subnetting Pattern for Scalable MLOps

This guide demonstrates a dynamic and scalable Terraform pattern using `cidrsubnet()` to create **public, private, and database subnets** across multiple Availability Zones (AZs). This is ideal for building robust, highly available MLOps or production cloud environments.

---

## ✅ 1. What You're Doing

- **Single VPC CIDR** (e.g., `10.0.0.0/16`)
- **Dynamic subnetting** (public/private/db) per AZ
- **No manual IP calculation** — everything is computed
- **Easily scales** with number of AZs and subnets

---

## 📘 2. How `cidrsubnet()` Works

Terraform’s [`cidrsubnet()`](https://developer.hashicorp.com/terraform/language/functions/cidrsubnet) function splits a base CIDR into non-overlapping subnets.

**Syntax:**
```
cidrsubnet(base_cidr_block, newbits, netnum)
```
- **base_cidr_block:** Your base network, e.g., `10.0.0.0/16`
- **newbits:** How many bits to add for new subnet mask (e.g., `8` means `/24`)
- **netnum:** The subnet index (0, 1, 2...)

**Example:**
```
cidrsubnet("10.0.0.0/16", 8, 0) → 10.0.0.0/24
cidrsubnet("10.0.0.0/16", 8, 1) → 10.0.1.0/24
```

---

## 📄 3. Terraform Example

### 🔧 Variables (`variables.tf`)
```hcl
variable "vpc_cidr" {
  default = "10.0.0.0/16"
}
variable "az_count" {
  default = 3
}
```

### 🗃️ Data Block for AZs
```hcl
data "aws_availability_zones" "available" {}
```

### 📁 Subnets (`main.tf`)
```hcl
# Public Subnets
resource "aws_subnet" "public" {
  count                   = var.az_count
  vpc_id                  = aws_vpc.main.id
  cidr_block              = cidrsubnet(var.vpc_cidr, 8, count.index)
  availability_zone       = data.aws_availability_zones.available.names[count.index]
  map_public_ip_on_launch = true
  tags = {
    Name = "public-${count.index}"
  }
}

# Private Subnets (shifted by az_count)
resource "aws_subnet" "private" {
  count             = var.az_count
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(var.vpc_cidr, 8, count.index + var.az_count)
  availability_zone = data.aws_availability_zones.available.names[count.index]
  tags = {
    Name = "private-${count.index}"
  }
}

# Database Subnets (shifted by 2 * az_count)
resource "aws_subnet" "database" {
  count             = var.az_count
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(var.vpc_cidr, 8, count.index + 2 * var.az_count)
  availability_zone = data.aws_availability_zones.available.names[count.index]
  tags = {
    Name = "db-${count.index}"
  }
}
```

### 🔄 Optional: Output Subnet Lists
```hcl
output "public_subnet_ids" {
  value = aws_subnet.public[*].id
}

output "private_subnet_ids" {
  value = aws_subnet.private[*].id
}

output "database_subnet_ids" {
  value = aws_subnet.database[*].id
}
```

---

## 🧠 4. Detailed Explanation

| Subnet Type | Index Formula                | Example CIDRs (for 3 AZs)             |
|-------------|----------------------------- |---------------------------------------|
| Public      | `count.index`                | 10.0.0.0/24, 10.0.1.0/24, 10.0.2.0/24 |
| Private     | `count.index + var.az_count` | 10.0.3.0/24, 10.0.4.0/24, 10.0.5.0/24 |
| Database    | `count.index + 2*az_count`   | 10.0.6.0/24, 10.0.7.0/24, 10.0.8.0/24 |

- Each subnet type uses a unique offset to avoid overlaps.
- Each AZ gets one public, one private, and one db subnet.

**Mapping Example:**
- **AZ0:** public-0, private-0, db-0
- **AZ1:** public-1, private-1, db-1
- **AZ2:** public-2, private-2, db-2

---

## 🚀 5. Why This Rocks (Best Practices)

- **Scalable:** Add/remove AZs by changing `az_count`
- **No IP Collisions:** Each tier gets non-overlapping space
- **Maintainable:** No hardcoded IPs, super readable
- **Multi-AZ Ready:** Ideal for HA, production, or MLOps
- **Automation-Friendly:** Integrates with modules, for_each, workspaces, etc.

---

## 🧪 Bonus: Turn This into a Module?

Want to make this a reusable module?
- Parameterize VPC CIDR, subnet mask, tiers, AZ filtering, etc.
- Consume as a module in multiple environments.

---

## 📚 References

- [Terraform: cidrsubnet()](https://developer.hashicorp.com/terraform/language/functions/cidrsubnet)
- [AWS VPC Best Practices](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html)
- [Terraform AWS Subnet Resource](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet)

---

**Tip:**  
This pattern is widely used for production-grade, multi-AZ, and MLOps-ready AWS architectures!
