Great question — **Google Cloud actually makes this easy**, and it’s one of the nicest backends Terraform supports 👍

### Short answer

👉 **Terraform uses Google Cloud Storage (GCS) object preconditions for locking**
👉 **No extra service, no DB, no add-ons**
👉 Locking is **native, automatic, and safe**

This is *not* block storage (Persistent Disk). Terraform state + locking on GCP is done with **GCS buckets**.

---

## How Terraform locking works on Google Cloud

When you use the **`gcs` backend**:

* Terraform writes the state file to GCS
* GCS provides **atomic object operations**
* Terraform uses **generation numbers** to ensure:

  * Only one writer at a time
  * Concurrent applies fail cleanly

💡 This gives you **real locking**, similar in strength to DynamoDB or Terraform Cloud.

---

## GCS backend example (with locking)

```hcl
terraform {
  backend "gcs" {
    bucket  = "my-terraform-state"
    prefix  = "prod"
  }
}
```

That’s it.
No extra flags. No lock tables. No coordination needed.

---

## What happens during a lock?

* Terraform starts an operation
* GCS enforces a **write precondition**
* If another run tries to write:

  * ❌ It fails with a lock error
* When the run finishes:

  * ✅ Lock is released automatically

If a process crashes, the next run can safely continue because GCS guarantees object consistency.

---

## IAM permissions required

Your service account needs:

```text
storage.objects.get
storage.objects.create
storage.objects.delete
storage.objects.update
```

On the bucket (or prefix).

---

## Why block storage (Persistent Disk) is NOT used

Terraform **does not support block storage backends** for state.

Reasons:

* No native locking API
* No atomic write guarantees
* Requires mounting disks everywhere Terraform runs

So:

* ❌ Persistent Disk
* ❌ Filestore
* ❌ Local disk + NFS

✅ **GCS is the correct and supported solution**

---

## Comparison (AWS vs GCP)

| Cloud | Backend  | Locking method            |
| ----- | -------- | ------------------------- |
| AWS   | S3 (new) | Native S3 lockfile        |
| AWS   | S3 (old) | DynamoDB                  |
| GCP   | GCS      | Object generation locking |
| Azure | Blob     | Lease locking             |

GCP’s approach is actually one of the cleanest.

---

## Best practices on GCP

* Enable **bucket versioning**
* Use **one bucket, multiple prefixes** per env
* Run Terraform via **CI with a single service account**
* Don’t mix local + CI applies casually

---

## TL;DR

* ✅ Terraform locking on GCP is **built-in**
* ✅ Uses **GCS**, not block storage
* ✅ No extra infra required
* ❌ Block storage is not supported or needed

If you want, I can show:

* Multi-env layout (`dev/stage/prod`)
* CI/CD example with GCP auth
* Migration from S3 → GCS state

Just tell me how you’re running Terraform today.
