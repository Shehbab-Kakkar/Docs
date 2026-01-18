To give **access to the `testOK-dev` GCS bucket only to Pods running in the `test-dev` GKE cluster**, the **correct and secure way** is to use **Workload Identity** scoped to that cluster.

Below is a **clean, least-privilege setup**. This ensures:

* ✅ Only Pods in **test-dev cluster**
* ✅ Only Pods using a **specific Kubernetes Service Account**
* ✅ Can access **only the `testOK-dev` bucket**

---

## 🔐 Architecture (what we’re doing)

```
GKE Pod
  ↓ (uses)
Kubernetes Service Account (KSA)
  ↓ (Workload Identity binding)
Google Service Account (GSA)
  ↓ (IAM role on bucket)
GCS bucket: testOK-dev
```

---

## 1️⃣ Ensure Workload Identity is enabled on `test-dev` cluster

```bash
gcloud container clusters describe test-dev \
  --format="value(workloadIdentityConfig.workloadPool)"
```

Expected output:

```
PROJECT_ID.svc.id.goog
```

If not enabled:

```bash
gcloud container clusters update test-dev \
  --workload-pool=PROJECT_ID.svc.id.goog
```

---

## 2️⃣ Create a **dedicated Google Service Account**

This GSA will have access **only to this bucket**.

```bash
gcloud iam service-accounts create testok-dev-gcs-sa \
  --display-name="GCS access for test-dev cluster"
```

---

## 3️⃣ Grant GCS access **ONLY on the bucket**

⚠️ Do **NOT** grant project-wide permissions.

### Read-only example

```bash
gsutil iam ch \
  serviceAccount:testok-dev-gcs-sa@PROJECT_ID.iam.gserviceaccount.com:roles/storage.objectViewer \
  gs://testOK-dev
```

### Read/write example

```bash
gsutil iam ch \
  serviceAccount:testok-dev-gcs-sa@PROJECT_ID.iam.gserviceaccount.com:roles/storage.objectAdmin \
  gs://testOK-dev
```

---

## 4️⃣ Create a Kubernetes Service Account (KSA)

Create it **only in the namespace you want** (recommended).

```bash
kubectl create namespace test-dev
kubectl create serviceaccount testok-gcs-ksa -n test-dev
```

---

## 5️⃣ Bind KSA → GSA (Workload Identity)

This step is what **locks access to this cluster’s pods only**.

```bash
gcloud iam service-accounts add-iam-policy-binding \
  testok-dev-gcs-sa@PROJECT_ID.iam.gserviceaccount.com \
  --role roles/iam.workloadIdentityUser \
  --member "serviceAccount:PROJECT_ID.svc.id.goog[test-dev/testok-gcs-ksa]"
```

---

## 6️⃣ Annotate the KSA

```bash
kubectl annotate serviceaccount testok-gcs-ksa \
  -n test-dev \
  iam.gke.io/gcp-service-account=testok-dev-gcs-sa@PROJECT_ID.iam.gserviceaccount.com
```

---
Below is a **complete, production-style `gcsfuse` example** for GKE that mounts the **`testOK-dev` GCS bucket** into a Pod at **`/data`**, using **Workload Identity** (no keys).

This works **only for Pods in your `test-dev` cluster + namespace** that use the bound KSA.

---

## ✅ Prerequisites (already done based on our earlier steps)

* GKE cluster: `test-dev`
* Workload Identity: **enabled**
* GSA: `testok-dev-gcs-sa`
* KSA: `testok-gcs-ksa` in namespace `test-dev`
* Bucket IAM:

  * `roles/storage.objectViewer` (read)
  * or `roles/storage.objectAdmin` (read/write)

---

## 1️⃣ Pod YAML – gcsfuse mount at `/data`

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: gcsfuse-testok
  namespace: test-dev
spec:
  serviceAccountName: testok-gcs-ksa
  restartPolicy: Never

  containers:
  - name: app
    image: gcr.io/google.com/cloudsdktool/cloud-sdk:slim
    command: ["/bin/bash", "-c"]
    args:
      - |
        echo "Listing files from GCS mount:"
        ls -lh /data
        echo "Copying files with 2026 in name to /tmp"
        cp /data/*2026* /tmp || true
        echo "Done. Sleeping..."
        sleep 3600
    volumeMounts:
    - name: gcs-bucket
      mountPath: /data

  volumes:
  - name: gcs-bucket
    csi:
      driver: gcsfuse.csi.storage.gke.io
      readOnly: false
      volumeAttributes:
        bucketName: testOK-dev
```

---

## 2️⃣ Deploy the Pod

```bash
kubectl apply -f gcsfuse-pod.yaml
```

---

## 3️⃣ Verify access

```bash
kubectl logs gcsfuse-testok -n test-dev
```

Exec into the pod:

```bash
kubectl exec -it gcsfuse-testok -n test-dev -- bash
ls -lh /data
```

---

## 📌 How this works (important)

* `gcsfuse.csi.storage.gke.io` mounts the bucket
* Uses **Workload Identity automatically**
* No service account keys
* IAM enforced at the **bucket level**
* Access limited to Pods using `testok-gcs-ksa`

---

## ⚠️ Notes & Best Practices

### Read-only mount (recommended if possible)

```yaml
readOnly: true
```

### Mount a subdirectory

```yaml
volumeAttributes:
  bucketName: testOK-dev
  mountOptions: "only-dir=reports/2026"
```

### Performance tuning

```yaml
mountOptions: "file-cache:max-size-mb:512,stat-cache-ttl:1m"
```

### When NOT to use gcsfuse

* Heavy random writes
* Databases
* High-throughput workloads

Use native GCS SDK instead in those cases.

---

## ✅ Security guarantees

✔ Only **test-dev cluster Pods**
✔ Only **testok-gcs-ksa**
✔ Only **testOK-dev bucket**
✔ No credentials stored

---

Here’s a **clear architecture diagram (text/ASCII)** showing how **GKE Pods in `test-dev`** securely access the **`testOK-dev` GCS bucket** using **Workload Identity + gcsfuse**.

---

## 🧭 High-Level Architecture Diagram

```
┌───────────────────────────────────────────────┐
│                 GKE CLUSTER                   │
│                test-dev (GKE)                 │
│                                               │
│  ┌─────────────────────────────────────────┐  │
│  │               Namespace                 │  │
│  │               test-dev                  │  │
│  │                                         │  │
│  │  ┌───────────────────────────────────┐  │  │
│  │  │               Pod                 │  │  │
│  │  │  name: gcsfuse-testok              │  │  │
│  │  │                                   │  │  │
│  │  │  /data  ───────────────┐          │  │  │
│  │  │        (gcsfuse mount)  │          │  │  │
│  │  │                         │          │  │  │
│  │  └─────────────────────────┘          │  │  │
│  │            ▲                           │  │  │
│  │            │                           │  │  │
│  │  ┌───────────────────────────────────┐  │  │
│  │  │ Kubernetes Service Account (KSA)  │  │  │
│  │  │  testok-gcs-ksa                   │  │  │
│  │  └───────────────────────────────────┘  │  │
│  │            ▲                           │  │  │
│  │            │ Workload Identity        │  │  │
│  │            ▼                           │  │  │
│  │  ┌───────────────────────────────────┐  │  │
│  │  │ Google Service Account (GSA)       │  │  │
│  │  │ testok-dev-gcs-sa                  │  │  │
│  │  └───────────────────────────────────┘  │  │
│  │            ▲                           │  │  │
│  │            │ IAM (bucket-level)       │  │  │
│  │            ▼                           │  │  │
│  └─────────────────────────────────────────┘  │
│                                               │
└───────────────────────────────────────────────┘
                │
                │ Google Cloud Storage API
                ▼
┌───────────────────────────────────────────────┐
│             GCS BUCKET                        │
│             gs://testOK-dev                   │
│                                               │
│  roles/storage.objectViewer / objectAdmin     │
│                                               │
└───────────────────────────────────────────────┘
```

---

## 🔁 Authentication & Access Flow (Step-by-Step)

```
Pod starts
   ↓
Uses KSA: testok-gcs-ksa
   ↓
Workload Identity exchange
   ↓
Acts as GSA: testok-dev-gcs-sa
   ↓
gcsfuse CSI driver
   ↓
GCS IAM checks (bucket-level)
   ↓
Bucket mounted at /data
```

---

## 🔐 Security Boundaries (why this is safe)

| Boundary                | Enforced by                |
| ----------------------- | -------------------------- |
| Only `test-dev` cluster | Workload Identity pool     |
| Only selected Pods      | Kubernetes Service Account |
| Only this bucket        | Bucket IAM                 |
| No credentials          | Metadata server            |
| No key leakage          | Zero JSON keys             |

---

## 🧩 Key Components Summary

| Component                 | Purpose          |
| ------------------------- | ---------------- |
| KSA (`testok-gcs-ksa`)    | Pod identity     |
| GSA (`testok-dev-gcs-sa`) | GCP IAM identity |
| Workload Identity         | Secure mapping   |
| gcsfuse CSI driver        | Filesystem mount |
| GCS bucket                | Storage backend  |

---

If you want, I can also:

* Convert this into a **Mermaid diagram** (for docs)
* Provide a **draw.io / Lucidchart layout**
* Show a **comparison diagram vs JSON keys**
* Add **multi-cluster isolation diagram**

Just tell me which format you want 👍
