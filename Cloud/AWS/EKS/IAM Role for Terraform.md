# 🚀 How to Create an IAM Role for Terraform to Assume (for EKS Creation)

This guide explains how to securely use an IAM role for Terraform to create and manage EKS clusters and related AWS resources, using the `sts:AssumeRole` mechanism.

---

## 1️⃣ Step-by-Step: Create an IAM Role for Terraform

### **Step 1: Create the IAM Role with an AssumeRole Trust Policy**

- The role must trust your Terraform execution entity (IAM user or role).
- Replace `<YOUR_TERRAFORM_EXECUTION_ENTITY_ARN>` with your actual user or role ARN.

**Example Trust Policy:**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::<YOUR_TERRAFORM_EXECUTION_ENTITY_ARN>"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
```
> _You can use the AWS Console, AWS CLI, or Terraform itself to create this role._

---

### **Step 2: Attach Required Policies**

Attach AWS managed (or custom) policies that allow Terraform to create EKS clusters and related resources.

**Recommended Managed Policies:**
- `AmazonEKSClusterPolicy`
- `AmazonEKSServicePolicy`
- `AmazonEKSWorkerNodePolicy`
- `AmazonEC2FullAccess` (or more restrictive EC2 permissions)
- `IAMFullAccess` (or scoped-down IAM permissions)
- `AmazonVPCFullAccess` (or scoped-down VPC permissions)

> _For production, scope these policies to the minimum permissions needed!_

---

### **Step 3: Configure Terraform to Assume the Role**

In your Terraform AWS provider configuration, specify the `role_arn` to assume:

```hcl
provider "aws" {
  region = "us-west-2"

  assume_role {
    role_arn = "arn:aws:iam::<ACCOUNT_ID>:role/TerraformEKSRole"
  }
}
```
- Replace `<ACCOUNT_ID>` and `TerraformEKSRole` with your values.

Terraform will use this role when making AWS API calls.

---

## 2️⃣ Why Use AssumeRole for Terraform?

| Aspect           | IAM Users / Service Accounts Directly | Using AssumeRole IAM Role for Terraform      |
|------------------|---------------------------------------|---------------------------------------------|
| **Security**     | Long-lived credentials, risk of leak  | Temporary credentials, reduced exposure     |
| **Separation**   | User has all privileges               | Terraform operates under scoped role        |
| **Access Mgmt**  | Harder to rotate/limit keys/permissions| Roles are easy to manage & rotate           |
| **Auditability** | Actions tied to raw user/service acct | Actions tied to role session (easier audit) |
| **Flexibility**  | Less flexible across accounts         | Can be assumed cross-account                |
| **Least Priv.**  | Often over-permissioned               | Easier to follow least privilege            |

---

## 3️⃣ Summary Example

### **A. Create the Role**
**Trust Policy Example:**
```json
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Principal": { "AWS": "arn:aws:iam::123456789012:user/terraform-user" },
    "Action": "sts:AssumeRole"
  }]
}
```
**Attach Managed Policies:**
- `AmazonEKSClusterPolicy`
- `AmazonEKSWorkerNodePolicy`
- `AmazonEC2FullAccess`
- `IAMFullAccess`
- (and others as needed)

### **B. Use in Terraform**
```hcl
provider "aws" {
  region = "us-west-2"
  assume_role {
    role_arn = "arn:aws:iam::123456789012:role/TerraformEKSRole"
  }
}
```

---

## ⭐ Best Practices

- Use **different roles** for different environments (dev, staging, prod)
- **Limit permissions** to only what's needed for EKS and resources
- **Review and rotate** trust relationships regularly
- **Enable MFA** on users allowed to assume the role for extra security

---

## 📚 References

- [AWS: Using IAM Roles with Terraform](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/guides/iam-assume-role)
- [AWS EKS IAM Requirements](https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html)
- [Terraform AWS Provider Docs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)

----
**IMPORTANT NOTE**

# 🛡️ Attaching Managed Policies to an IAM Role for Terraform EKS Automation

This guide shows how to attach AWS managed policies to an IAM Role for Terraform automation—using both the AWS CLI and Terraform.

---

## 1️⃣ Attach Managed Policies Using AWS CLI

**Example:** Attach `AmazonEKSClusterPolicy` to your role `TerraformEKSRole`:

```bash
aws iam attach-role-policy \
  --role-name TerraformEKSRole \
  --policy-arn arn:aws:iam::aws:policy/AmazonEKSClusterPolicy
```

**Repeat for each required policy:**

```bash
aws iam attach-role-policy --role-name TerraformEKSRole --policy-arn arn:aws:iam::aws:policy/AmazonEKSClusterPolicy
aws iam attach-role-policy --role-name TerraformEKSRole --policy-arn arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy
aws iam attach-role-policy --role-name TerraformEKSRole --policy-arn arn:aws:iam::aws:policy/AmazonEC2FullAccess
aws iam attach-role-policy --role-name TerraformEKSRole --policy-arn arn:aws:iam::aws:policy/IAMFullAccess
aws iam attach-role-policy --role-name TerraformEKSRole --policy-arn arn:aws:iam::aws:policy/AmazonVPCFullAccess
```

---

## 2️⃣ Attach Managed Policies Using Terraform

Define your IAM role and attach managed policies in Terraform as shown below:

```hcl
resource "aws_iam_role" "terraform_eks_role" {
  name = "TerraformEKSRole"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Principal = {
        AWS = "arn:aws:iam::123456789012:user/terraform-user"
      }
      Action = "sts:AssumeRole"
    }]
  })
}

resource "aws_iam_role_policy_attachment" "eks_cluster_policy" {
  role       = aws_iam_role.terraform_eks_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
}

resource "aws_iam_role_policy_attachment" "eks_worker_node_policy" {
  role       = aws_iam_role.terraform_eks_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
}

resource "aws_iam_role_policy_attachment" "ec2_full_access" {
  role       = aws_iam_role.terraform_eks_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2FullAccess"
}

resource "aws_iam_role_policy_attachment" "iam_full_access" {
  role       = aws_iam_role.terraform_eks_role.name
  policy_arn = "arn:aws:iam::aws:policy/IAMFullAccess"
}

resource "aws_iam_role_policy_attachment" "vpc_full_access" {
  role       = aws_iam_role.terraform_eks_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonVPCFullAccess"
}
```

---

### 📝 Explanation

- `aws_iam_role` creates the IAM Role with the specified trust (assume role) policy.
- Each `aws_iam_role_policy_attachment` resource attaches a managed policy to the role.
- This makes the role ready for Terraform to assume and manage EKS and related resources.

---

## 🔒 Security & Best Practices

- **Principle of Least Privilege:** Only attach policies your automation actually needs. Scope down from `*FullAccess` when possible.
- **Separate Roles for Environments:** Use distinct roles for dev/staging/prod.
- **Audit Regularly:** Check what policies are attached and review access.
- **MFA Requirement:** Require MFA for users who can assume the role.

---

## 📚 References

- [Terraform AWS IAM Role Policy Attachment Docs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment)
- [AWS CLI: attach-role-policy](https://docs.aws.amazon.com/cli/latest/reference/iam/attach-role-policy.html)
- [EKS IAM Requirements](https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html)

---

**Tip:**  
For production, always minimize permissions and monitor role usage!

---

**Security Tip:**  
Never use long-lived Access Keys for Terraform automation – always use short-lived AssumeRole credentials!
