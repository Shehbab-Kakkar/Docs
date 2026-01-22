

## 📦 Example Backend Configuration (no DynamoDB)

Here’s what your Terraform backend looks like with **native S3 locking**:

```hcl
terraform {
  backend "s3" {
    bucket        = "my-terraform-state"
    key           = "envs/prod/terraform.tfstate"
    region        = "ap-south-1"
    encrypt       = true
    use_lockfile  = true   # ← enables native S3 locking
  }
}
```

Then reinitialize:

```bash
terraform init -migrate-state
```

Now Terraform will use the S3 lockfile instead of DynamoDB for protecting the state. ([projectdevops.in][1])

---

## 🛠️ IAM Permissions Needed

To support this native locking, ensure the IAM role or user has the following on the state bucket:

```json
{
  "Effect": "Allow",
  "Action": [
    "s3:GetObject",
    "s3:PutObject",
    "s3:DeleteObject"
  ],
  "Resource": "arn:aws:s3:::my-terraform-state/*"
}
```

The lockfile operations are just S3 object operations — no DynamoDB permissions needed. ([projectdevops.in][1])

---

## 📌 Notes & Tips

* **Bucket versioning** is still recommended to help with state rollbacks. ([HashiCorp Developer][2])

* If a lockfile gets left behind (e.g., due to a crash), use:

  ```bash
  terraform force-unlock <LOCK_ID>
  ```

  just like with the old DynamoDB locks. ([HashiCorp Developer][3])

* DynamoDB locking is now **deprecated** and will be removed in a future version — migrating to native S3 locking early is a good idea. ([blog.dnyaneshwarsonawane.com][4])

---

## ✅ Summary

✔ You *can* have **state locking without DynamoDB**
✔ Terraform now supports **native S3 lockfiles**
✔ Just enable `use_lockfile = true` in your backend
✔ Keeps your setup simpler and cheaper

---

If you want, I can also show you the **exact Terraform version requirement** for this feature or how to **migrate from a DynamoDB setup** smoothly!

[1]: https://www.projectdevops.in/b/68cf9b9eda21c240cfb9fa4f-terraform-s3-lockfile-the-new-way-to-lock-state-without?utm_source=chatgpt.com "Terraform S3 Lockfile: The New Way to Lock State Without DynamoDB"
[2]: https://developer.hashicorp.com/terraform/language/backend/s3?utm_source=chatgpt.com "Backend Type: s3 | Terraform | HashiCorp Developer"
[3]: https://developer.hashicorp.com/terraform/language/state/locking?utm_source=chatgpt.com "State: Locking | Terraform | HashiCorp Developer"
[4]: https://blog.dnyaneshwarsonawane.com/p/terraform-110-native-s3-locking-no?utm_source=chatgpt.com "Terraform 1.10+: Native S3 Locking — No More DynamoDB Needed"
