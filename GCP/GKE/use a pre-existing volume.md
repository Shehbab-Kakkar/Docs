To **use a pre-existing volume in Google Kubernetes Engine (GKE)**, such as a **pre-created Persistent Disk (PD)** in Google Cloud, you’ll need to create a **PersistentVolume (PV)** and a **PersistentVolumeClaim (PVC)** that maps to it.

Here’s a **step-by-step guide**:

---

### ✅ Step 1: Identify Your Existing Persistent Disk

First, make sure you have a **Google Cloud Persistent Disk** created. You can do this via the Cloud Console or using the CLI:

```bash
gcloud compute disks list
```

Take note of:

* **Disk name**
* **Zone**
* **Size**
* **Type (pd-standard, pd-ssd, etc.)**

---

### ✅ Step 2: Create a PersistentVolume (PV)

This will represent the existing disk in Kubernetes.

```yaml
# pv.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: my-preexisting-pv
spec:
  capacity:
    storage: 10Gi  # Must match or be less than actual disk size
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  gcePersistentDisk:
    pdName: your-existing-disk-name
    fsType: ext4  # Or the actual filesystem type
  storageClassName: ""
```

> 📝 Note: Use `storageClassName: ""` to indicate that it’s a **static volume** not managed by dynamic provisioning.

Apply it:

```bash
kubectl apply -f pv.yaml
```

---

### ✅ Step 3: Create a PersistentVolumeClaim (PVC)

This is what your pods will use to request the storage.

```yaml
# pvc.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
  volumeName: my-preexisting-pv
  storageClassName: ""
```

Apply it:

```bash
kubectl apply -f pvc.yaml
```

---

### ✅ Step 4: Use the PVC in a Pod or Deployment

Example pod:

```yaml
# pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
    - name: my-container
      image: nginx
      volumeMounts:
        - mountPath: "/mnt/data"
          name: my-storage
  volumes:
    - name: my-storage
      persistentVolumeClaim:
        claimName: my-pvc
```

Apply it:

```bash
kubectl apply -f pod.yaml
```

---

### 🧹 Optional: Disk Reuse and Cleanup Tips

* Use `persistentVolumeReclaimPolicy: Retain` to prevent GKE from deleting your disk after PVC is deleted.
* To reuse the PV with a new PVC, make sure to **manually delete the old PVC or PV**, or adjust the `volumeName` and `claimRef`.

---

### 🔎 Troubleshooting

* If the pod is stuck in `Pending`, check:

  * PV and PVC names match
  * `storageClassName` is set to `""` in both
  * Disk is in the same zone as the node/pod
* Use:

```bash
kubectl describe pvc my-pvc
kubectl describe pv my-preexisting-pv
kubectl get events
```

---

If you’d like, I can generate these YAML files based on your disk name and details. Just share:

* Disk name
* Disk size
* Filesystem (ext4, xfs, etc.)
* Access mode needed (RWO, ROX, RWX)
* GKE cluster zone

Let me know.
