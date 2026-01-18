Here’s a clear, practical overview of **how a Pod in GKE can access Google Cloud Storage (GCS)**. I’ll cover the **recommended approach first**, then alternatives, with examples.

---

## ✅ Recommended: **Workload Identity (no keys)**

**Workload Identity** lets a Kubernetes Service Account (KSA) act as a Google Service Account (GSA) securely—**no JSON keys in pods**.

### 1️⃣ Enable Workload Identity on the cluster

```bash
gcloud container clusters update CLUSTER_NAME \
  --workload-pool=PROJECT_ID.svc.id.goog
```

---

### 2️⃣ Create a Google Service Account

```bash
gcloud iam service-accounts create gcs-access-sa
```

Grant it GCS permissions:

```bash
gcloud projects add-iam-policy-binding PROJECT_ID \
  --member="serviceAccount:gcs-access-sa@PROJECT_ID.iam.gserviceaccount.com" \
  --role="roles/storage.objectViewer"
```

(or `objectAdmin` if you need write access)

---

### 3️⃣ Create a Kubernetes Service Account

```bash
kubectl create serviceaccount gcs-access-ksa
```

---

### 4️⃣ Bind KSA → GSA

```bash
gcloud iam service-accounts add-iam-policy-binding \
  gcs-access-sa@PROJECT_ID.iam.gserviceaccount.com \
  --member="serviceAccount:PROJECT_ID.svc.id.goog[default/gcs-access-ksa]" \
  --role="roles/iam.workloadIdentityUser"
```

Annotate the KSA:

```bash
kubectl annotate serviceaccount gcs-access-ksa \
  iam.gke.io/gcp-service-account=gcs-access-sa@PROJECT_ID.iam.gserviceaccount.com
```

---

### 5️⃣ Use the KSA in your Pod

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: app
spec:
  serviceAccountName: gcs-access-ksa
  containers:
  - name: app
    image: gcr.io/my-project/my-app
```

Now your app can access GCS using **Google Cloud SDKs** automatically.

---

### 6️⃣ Access GCS from code (example: Python)

```python
from google.cloud import storage

client = storage.Client()
bucket = client.bucket("my-bucket")
blob = bucket.blob("file.txt")
print(blob.download_as_text())
```

No credentials needed 🎉

---

## 🔁 Alternative Options

### Option A: **gcsfuse (mount GCS as filesystem)**

Useful for legacy apps that expect files.

```yaml
containers:
- name: app
  image: my-app
  volumeMounts:
  - mountPath: /data
    name: gcs-volume
volumes:
- name: gcs-volume
  csi:
    driver: gcsfuse.csi.storage.gke.io
    readOnly: false
```

✔ Uses Workload Identity
⚠ Slower than native GCS API for heavy I/O

---

### Option B: **Service Account JSON key (NOT recommended)**

* Store key in Kubernetes Secret
* Mount into pod
* Set `GOOGLE_APPLICATION_CREDENTIALS`

❌ Security risk
❌ Key rotation overhead
Use only if Workload Identity is impossible

---

## 🧭 Which should you use?

| Use case                 | Best option               |
| ------------------------ | ------------------------- |
| Standard apps            | ✅ Workload Identity + SDK |
| Needs filesystem access  | gcsfuse                   |
| Legacy / non-GCP cluster | JSON key (last resort)    |

---

If you want, tell me:

* **Language** (Python, Java, Node, Go)
* **Read vs write**
* **SDK vs filesystem**
* **Autopilot or Standard GKE**

I can give you a **copy-paste minimal example** tailored to your setup.
