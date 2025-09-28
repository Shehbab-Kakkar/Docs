In Google Cloud Platform (GCP), **permissions** control access to resources. Permissions are bundled into **roles**, and these roles are granted to **users**, **groups**, or **service accounts** to perform actions on GCP resources.

---

### 🔑 Basics of GCP Permissions

* **Permission**: Allows a specific action, like `compute.instances.get`.
* **Role**: A collection of permissions (e.g., `roles/viewer` includes read-only permissions).
* **IAM Policy**: Binds a role to a member (user/service account) on a resource.

---

## ✅ Example Scenario

> **Goal**: Grant a user permission to start and stop Compute Engine VM instances.

You’ll need to assign the `Compute Instance Admin (v1)` role to the user, which includes permissions like:

* `compute.instances.start`
* `compute.instances.stop`
* `compute.instances.get`

---

## 🔧 Example Command: Grant IAM Role to a User

```bash
gcloud projects add-iam-policy-binding <PROJECT_ID> \
  --member="user:example@gmail.com" \
  --role="roles/compute.instanceAdmin.v1"
```

### 🔍 Explanation:

| Part                                      | Description                                                    |
| ----------------------------------------- | -------------------------------------------------------------- |
| `gcloud projects add-iam-policy-binding`  | Adds a role binding to the project’s IAM policy                |
| `<PROJECT_ID>`                            | Your GCP project ID (e.g., `my-gcp-project`)                   |
| `--member="user:example@gmail.com"`       | The member (user) receiving the role                           |
| `--role="roles/compute.instanceAdmin.v1"` | The role that includes necessary permissions for VM management |

---

## 🔒 Verifying Permissions

To **check what permissions** a role has:

```bash
gcloud iam roles describe roles/compute.instanceAdmin.v1
```

This shows the list of permissions included in that role.

---

## 🧪 Testing Permissions (Optional)

You can simulate a permission test:

```bash
gcloud iam testable-permissions //compute.googleapis.com/projects/<PROJECT_ID>/zones/<ZONE>/instances/<INSTANCE_NAME>
```

Or use:

```bash
gcloud projects get-iam-policy <PROJECT_ID>
```

To view the entire IAM policy of the project.

---
