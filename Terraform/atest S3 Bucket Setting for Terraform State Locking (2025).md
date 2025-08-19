text
# Terraform S3 Backend — Latest Native Locking Setup (2025)

Terraform **v1.10+** now supports **native S3 backend locking**, letting you manage remote state without DynamoDB.  
This guide covers the recommended S3 bucket settings for secure, reliable remote state management.

---

## 🛠️ S3 Bucket Recommended Settings
- **Enable Versioning** → Protects state from accidental deletion or corruption.  
- **Server-Side Encryption** → Keep state files secure at rest.  
- **ACL: private** → Restricts bucket access to authorized IAM users.  
- **Native S3 Locking** → Use Terraform’s `use_lockfile = true` for lock management — **no DynamoDB required!**  
- **Minimum Terraform Version** → `>= 1.10`

---

## ⚡ Example: Backend Configuration (HCL)

terraform {
backend "s3" {
bucket = "your-terraform-state-bucket"
key = "env/terraform.tfstate"
region = "us-east-1"
encrypt = true
use_lockfile = true # Native S3 state locking!
acl = "private"
workspace_key_prefix = "env:"
}
}

text

---

## 🚀 S3 Bucket Setup (CLI Snippets)

**Create the S3 bucket**
aws s3api create-bucket --bucket <your-bucket-name> --region <region>

text

**Enable versioning for backup & recovery**
aws s3api put-bucket-versioning --bucket <your-bucket-name>
--versioning-configuration Status=Enabled

text

**Enable server-side encryption (recommended)**
aws s3api put-bucket-encryption --bucket <your-bucket-name>
--server-side-encryption-configuration '{
"Rules":[{"ApplyServerSideEncryptionByDefault":{"SSEAlgorithm":"AES256"}}]
}'

text

---

## 👍 Best Practices
- Use **dedicated S3 buckets** for Terraform state storage.  
- Store and access **AWS credentials securely**.  
- Regularly **audit S3 bucket access logs**.  
- **Enable versioning and encryption — always.**  
- Upgrade **Terraform to v1.10+** to use `use_lockfile`.  

---

## 🎯 Key Points
- Native **S3 state locking** is built-in → just set `use_lockfile = true`.  
- **DynamoDB table setup is NOT needed.**  
- Always **secure & version** your S3 bucket!  

---

> 📖 See [Terraform Docs on S3 Backend](https://developer.hashicorp.com/terraform/language/backend/s3) for additional details.
