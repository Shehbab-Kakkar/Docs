
### Terraform S3 State Locking Without DynamoDB

Terraform's S3 backend traditionally used DynamoDB for state locking to prevent concurrent modifications and ensure consistency. However, recent updates (starting from Terraform 1.10) introduce native S3 locking as an experimental feature, eliminating the need for DynamoDB. This simplifies your setup by relying solely on S3's conditional writes and object metadata for concurrency control.[^1_1][^1_2][^1_3][^1_4]

This approach creates a temporary `.tflock` file in your S3 bucket during operations like `apply` or `destroy`, which acts as the lock. If another process tries to modify the state simultaneously, it fails due to S3's checks, displaying an "Error acquiring the state lock" message.[^1_3][^1_5][^1_4][^1_1]

#### Key Prerequisites

- Use Terraform version 1.10 or higher (ideally 1.11+ for stability).[^1_2][^1_3]
- Enable versioning on your S3 bucket to support state file management.[^1_5][^1_1]
- Optionally enable object lock on the bucket for added consistency, though it's not strictly required.[^1_5]
- Ensure your IAM permissions allow S3 read/write access, including conditional operations.[^1_2]


#### Enabling Native S3 Locking

To switch, update your backend configuration by removing the `dynamodb_table` attribute and adding `use_lockfile = true`. This is optional for now and defaults to `false`, but it enables S3-native locking.[^1_1][^1_3][^1_5]

Here's a before-and-after example:

**Previous Configuration (With DynamoDB)**

```hcl
terraform {
  backend "s3" {
    bucket         = "your-terraform-state-bucket"
    key            = "path/to/your/statefile.tfstate"
    region         = "us-east-1"
    dynamodb_table = "your-dynamodb-lock-table"  # Required for locking
    encrypt        = true
  }
}
```

**Updated Configuration (Without DynamoDB)**

```hcl
terraform {
  backend "s3" {
    bucket       = "your-terraform-state-bucket"
    key          = "path/to/your/statefile.tfstate"
    region       = "us-east-1"
    encrypt      = true
    use_lockfile = true  # Enables S3 native locking
  }
}
```

After updating, run `terraform init` to reconfigure the backend. During operations, you'll see locking in action—e.g., a concurrent `terraform apply` will fail with a lock error.[^1_4][^1_1]

#### Example: Creating an S3 Bucket with Native Locking

For a full setup, you might define resources like this while using the backend:[^1_1]

```hcl
terraform {
  backend "s3" {
    bucket       = "ct-bucket-2025"
    key          = "cloudthat/terraform.tfstate"
    region       = "us-east-1"
    encrypt      = true
    use_lockfile = true  # S3 native locking
  }
}

resource "aws_s3_bucket" "ct_dev_bucket" {
  bucket = "ct-terraform-state-bucket"
  acl    = "private"
}

resource "aws_s3_bucket_versioning" "example_versioning" {
  bucket = aws_s3_bucket.ct_dev_bucket.id
  versioning_configuration {
    status = "Enabled"
  }
}
```

Apply this with `terraform apply`, and check your S3 bucket for the state file and temporary `.tflock` during execution.[^1_3][^1_1]

#### Benefits and Limitations

- **Benefits**: Reduces costs and complexity by avoiding DynamoDB setup and billing. It's ideal for simpler workflows or smaller teams with infrequent Terraform runs.[^1_6][^1_4][^1_5][^1_1]
- **Limitations**: As an experimental feature, it relies on S3's eventual consistency, which might lead to rare race conditions in high-concurrency environments. DynamoDB offers stronger guarantees for large-scale or frequent operations. In future Terraform releases, DynamoDB support may be deprecated in favor of this method.[^1_7][^1_4][^1_3]

If your workflow involves multiple users or high activity, consider sticking with DynamoDB for now. For alternatives, OpenTofu (a Terraform fork) explores S3-only locking via a `lock_storage_type` option, but it's not standard in Terraform.[^1_6][^1_4]

This setup should cover most use cases—test it in a non-production environment first to ensure it fits your needs.
<span style="display:none">[^1_10][^1_8][^1_9]</span>

<div style="text-align: center">⁂</div>

[^1_1]: https://www.cloudthat.com/resources/blog/state-locking-with-s3-without-dynamodb-in-terraform/

[^1_2]: https://www.youtube.com/watch?v=jGIWlA9d0Zg

[^1_3]: https://dev.to/globart/s3-native-state-locking-in-terraform-518i

[^1_4]: https://blog.techiescamp.com/terraform-state-locking/

[^1_5]: https://www.linkedin.com/pulse/terraform-remote-backend-configuration-s3-without-dynamodb-joish-ju1je

[^1_6]: https://github.com/opentofu/opentofu/issues/599

[^1_7]: https://developer.hashicorp.com/terraform/language/backend/s3

[^1_8]: https://discuss.hashicorp.com/t/feature-request-terraform-state-locking-in-aws-with-s3-strong-consistency-no-dynamodb/18456

[^1_9]: https://aws.plainenglish.io/farewell-to-dynamodb-terraform-s3-backend-supports-native-locking-4da4dbba4ef6

[^1_10]: https://github.com/hashicorp/terraform/issues/35625


---

# Last feature I want to use where I do locking

The **last feature you used for locking** in Terraform was the native S3 backend state locking introduced in Terraform **v1.10+**. This feature lets you enable locking without DynamoDB by setting `use_lockfile = true` in your backend configuration.

Key aspects:

- **Native S3 state locking** with `use_lockfile = true`
- No DynamoDB table setup required
- Best used with S3 versioning and encryption enabled
- Minimum Terraform version required: `>= 1.10`[^2_1]

Example configuration:

```hcl
terraform {
  backend "s3" {
    bucket         = "your-terraform-state-bucket"
    key            = "env/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    use_lockfile   = true # Native S3 state locking!
    acl            = "private"
    workspace_key_prefix = "env:"
  }
}
```

**Summary**: Your last locking feature was S3-native state locking using the `use_lockfile` option, completely eliminating the need for DynamoDB setup for state management.
<span style="display:none">[^2_2][^2_3][^2_4]</span>

<div style="text-align: center">⁂</div>

[^2_1]: https://github.com/Shehbab-Kakkar/Docs/blob/main/Terraform/atest S3 Bucket Setting for Terraform State Locking (2025).md

[^2_2]: comet://newtab/

[^2_3]: https://github.com/Shehbab-Kakkar/Docs/blob/main/Terraform/atest S3 Bucket Setting for Terraform State Locking (2025).md

[^2_4]: https://github.com/Shehbab-Kakkar/Docs/edit/main/Terraform/atest S3 Bucket Setting for Terraform State Locking (2025).md

