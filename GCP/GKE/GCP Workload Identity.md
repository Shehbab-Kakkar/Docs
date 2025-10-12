**GCP Workload Identity** is a way to allow applications running in **Kubernetes (GKE)** to access **Google Cloud resources securely**—without needing to manage service account keys.

Instead of using service account keys (which can be a security risk), GCP Workload Identity allows a Kubernetes service account to **impersonate** a Google Cloud service account.

---

### 🔹 Real-World Example of GCP Workload Identity

**Scenario**:
You have an application running in a GKE cluster that needs to access a **Google Cloud Storage (GCS)** bucket to read some files.

---

### 🔧 Steps (with Example Names)

#### 1. **Google Cloud Service Account (GSA)**

Create a GCP service account that has access to the GCS bucket.

```bash
gcloud iam service-accounts create gcs-reader \
  --project=my-gcp-project
```

Give it permission to read from the bucket:

```bash
gcloud projects add-iam-policy-binding my-gcp-project \
  --member="serviceAccount:gcs-reader@my-gcp-project.iam.gserviceaccount.com" \
  --role="roles/storage.objectViewer"
```

---

#### 2. **Kubernetes Service Account (KSA)**

Create a Kubernetes service account in your GKE cluster:

```bash
kubectl create serviceaccount k8s-gcs-reader
```

---

#### 3. **Bind KSA to GSA (Workload Identity Binding)**

Allow the Kubernetes service account (`k8s-gcs-reader`) to impersonate the Google service account (`gcs-reader@my-gcp-project.iam.gserviceaccount.com`):

```bash
gcloud iam service-accounts add-iam-policy-binding gcs-reader@my-gcp-project.iam.gserviceaccount.com \
  --role="roles/iam.workloadIdentityUser" \
  --member="serviceAccount:my-gcp-project.svc.id.goog[default/k8s-gcs-reader]"
```

> 🔹 Note: Replace `default` with your namespace if not using the default namespace.

---

#### 4. **Annotate the Kubernetes Service Account**

Tell GKE which Google service account the KSA should impersonate:

```bash
kubectl annotate serviceaccount k8s-gcs-reader \
  iam.gke.io/gcp-service-account=gcs-reader@my-gcp-project.iam.gserviceaccount.com
```

---

#### 5. **Deploy the Pod with the KSA**

Deploy your workload (e.g., a Pod) using the Kubernetes service account:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: gcs-app
spec:
  serviceAccountName: k8s-gcs-reader
  containers:
  - name: app
    image: google/cloud-sdk
    command: ["sleep", "3600"]
```

Once running, your pod can access GCS using the **Google service account's identity**, **without needing a key**.

---

### ✅ Benefits

* **No service account keys** to manage or rotate.
* **Fine-grained access** using IAM roles.
* **Secure and scalable** identity for GKE workloads.

---

g
