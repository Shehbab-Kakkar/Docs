Taking a volume snapshot in **Google Kubernetes Engine (GKE)** involves using **Kubernetes VolumeSnapshot** resources provided by the **Kubernetes CSI Snapshot** feature. This is commonly used to back up or preserve the state of a PersistentVolumeClaim (PVC).

Below is a **complete example** showing how to create a volume snapshot in GKE, assuming you're using a CSI-compatible storage class (like `pd-ssd` or `pd-standard` with the **GCP Persistent Disk CSI driver**).

---

### ✅ Prerequisites

1. GKE cluster with CSI driver enabled (`gce-pd-csi-driver`)
2. Kubernetes version 1.17+
3. Snapshot CRDs installed (`VolumeSnapshot`, `VolumeSnapshotContent`, and `VolumeSnapshotClass`)
4. PVC already created and bound

---

### 🔧 Step 1: Install the Snapshot CRDs and Controller (if not already present)

GKE Autopilot often includes this by default. For Standard clusters:

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/v6.2.1/client/config/crd/snapshot.storage.k8s.io_volumesnapshotclasses.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/v6.2.1/client/config/crd/snapshot.storage.k8s.io_volumesnapshotcontents.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/v6.2.1/client/config/crd/snapshot.storage.k8s.io_volumesnapshots.yaml
```

Install the snapshot controller (only needed in Standard mode, not Autopilot):

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/v6.2.1/deploy/kubernetes/snapshot-controller/rbac-snapshot-controller.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/v6.2.1/deploy/kubernetes/snapshot-controller/setup-snapshot-controller.yaml
```

---

### 📁 Step 2: Create a `VolumeSnapshotClass`

This defines how snapshots are created.

```yaml
# volumesnapshotclass.yaml
apiVersion: snapshot.storage.k8s.io/v1
kind: VolumeSnapshotClass
metadata:
  name: gcp-pd-snapshot-class
driver: pd.csi.storage.gke.io
deletionPolicy: Delete  # Or Retain
```

Apply it:

```bash
kubectl apply -f volumesnapshotclass.yaml
```

---

### 📦 Step 3: Create a PersistentVolumeClaim (PVC)

If you don’t already have one:

```yaml
# pvc.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: demo-pvc
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: standard-rwo
  resources:
    requests:
      storage: 10Gi
```

Apply it:

```bash
kubectl apply -f pvc.yaml
```

Wait for the PVC to be bound:

```bash
kubectl get pvc demo-pvc
```

---

### 📸 Step 4: Create a `VolumeSnapshot`

```yaml
# snapshot.yaml
apiVersion: snapshot.storage.k8s.io/v1
kind: VolumeSnapshot
metadata:
  name: demo-pvc-snapshot
spec:
  volumeSnapshotClassName: gcp-pd-snapshot-class
  source:
    persistentVolumeClaimName: demo-pvc
```

Apply it:

```bash
kubectl apply -f snapshot.yaml
```

---

### ✅ Step 5: Verify the Snapshot

Check the snapshot status:

```bash
kubectl get volumesnapshot
kubectl describe volumesnapshot demo-pvc-snapshot
```

Wait until `ReadyToUse: true`.

---

### 🧪 Bonus: Restore from the Snapshot

To create a new PVC from a snapshot:

```yaml
# pvc-from-snapshot.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: restored-pvc
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: standard-rwo
  resources:
    requests:
      storage: 10Gi
  dataSource:
    name: demo-pvc-snapshot
    kind: VolumeSnapshot
    apiGroup: snapshot.storage.k8s.io
```

Apply it:

```bash
kubectl apply -f pvc-from-snapshot.yaml
```

---

