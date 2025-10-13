To **completely delete a GKE (Google Kubernetes Engine) cluster along with its worker nodes**, you can use the **Google Cloud Console**, **gcloud CLI**, or **Terraform** (if you used it to create the cluster).

Here's how to do it using each method:

---

## ✅ Using `gcloud` CLI (Recommended)

This is the most direct and complete way to delete everything related to a GKE cluster.

### 1. **Set your project and zone/region**

```bash
gcloud config set project [PROJECT_ID]
gcloud config set compute/zone [COMPUTE_ZONE]  # for zonal clusters
# or
gcloud config set compute/region [COMPUTE_REGION]  # for regional clusters
```

### 2. **List GKE clusters (optional)**

```bash
gcloud container clusters list
```

### 3. **Delete the GKE cluster**

```bash
gcloud container clusters delete [CLUSTER_NAME]
```

* This command **deletes the control plane and all associated node pools (worker nodes)**.
* If you’re using **auto-provisioning**, this will also remove automatically created node pools.

You will be prompted for confirmation. Add `--quiet` to skip prompts:

```bash
gcloud container clusters delete [CLUSTER_NAME] --quiet
```

---

## 📍 Notes

* Deleting a GKE cluster will:

  * Delete the **control plane**
  * Delete all **node pools** (worker nodes)
  * Delete associated **VM instances**
  * **Detach and delete** persistent disks (unless you've configured them to persist)
  * **Remove firewall rules**, routes, and other GCP resources created for the cluster

* If you manually created external resources (like Cloud SQL, VPC, or persistent disks), you'll need to clean those up separately.

---

## 🖥️ Using Google Cloud Console

1. Go to **Google Kubernetes Engine > Clusters**.
2. Find the cluster you want to delete.
3. Click the **trash can** icon or "Delete" option.
4. Confirm deletion.

This will remove everything just like the CLI.

---

## 🛠️ Using Terraform

If you created the cluster using Terraform, use:

```bash
terraform destroy
```

Ensure your Terraform state is up to date and targeting the right resources/project.

---

## ✅ Final Check

After deletion, you can verify everything is gone:

* Check GKE clusters:

  ```bash
  gcloud container clusters list
  ```

* Check Compute Engine for remaining VMs:

  ```bash
  gcloud compute instances list
  ```

* Check for leftover disks:

  ```bash
  gcloud compute disks list
  ```

---

Let me know if you’re using **regional clusters**, **private clusters**, or **Terraform**, and I can give you more tailored instructions.
