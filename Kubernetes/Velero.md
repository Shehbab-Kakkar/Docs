Restoring a **Kubernetes cluster** with a **corrupted etcd** when using **Verelio** (assuming you're referring to Velero—used for backup and restore of Kubernetes objects) involves a series of careful steps. Since `etcd` is the source of truth for cluster state, corruption means you’ll need to restore it from a backup.

Here’s a **step-by-step guide** to restore a Kubernetes cluster if `etcd` is corrupted using **Velero** and/or `etcd` snapshots:

---

### 🔍 First: Understand What’s Backed Up

* **etcd** stores the actual cluster state: pods, services, config maps, secrets, deployments, etc.
* **Velero** backs up Kubernetes resources by querying the API server, not the etcd DB directly.
* So if `etcd` is **completely lost/corrupt**, **Velero alone is not enough** unless the cluster can come back online to apply the Velero backup.

If you’ve lost etcd and can’t start the API server, you need to:

1. **Restore etcd first (if you have a snapshot)**
2. **Then restore Velero backup (if needed) for higher-layer resources**

---

## 🔁 Option 1: Restore from etcd Snapshot (Preferred)

If you have an etcd snapshot from before the corruption:

### 1. Stop the Kubernetes Control Plane

```bash
sudo systemctl stop kube-apiserver
```

### 2. Move or Delete the Corrupted etcd Data

```bash
sudo mv /var/lib/etcd /var/lib/etcd.bak
```

### 3. Restore etcd from Snapshot

```bash
ETCDCTL_API=3 etcdctl snapshot restore /path/to/etcd-snapshot.db \
  --data-dir /var/lib/etcd
```

### 4. Adjust ownership/permissions if needed

```bash
sudo chown -R etcd:etcd /var/lib/etcd
```

### 5. Restart etcd and kube-apiserver

```bash
sudo systemctl start etcd
sudo systemctl start kube-apiserver
```

---

## 🔁 Option 2: Rebuild Cluster and Use Velero to Restore

If you **don’t** have an etcd snapshot:

You’ll need to **recreate the control plane** (basically reinstall the cluster) and **restore Kubernetes resources** from **Velero** backups.

### Step-by-step:

#### 1. Recreate the Cluster

Rebuild your control plane nodes. You can use:

* `kubeadm` to reinitialize the cluster
* Or redeploy using your provisioning tool (kOps, RKE, etc.)

Example with `kubeadm`:

```bash
kubeadm init --config kubeadm-config.yaml
```

#### 2. Reinstall Velero

Make sure Velero is installed in the new cluster and connected to your backup storage (e.g., S3, Azure Blob, etc.)

#### 3. Restore Velero Backup

List available backups:

```bash
velero backup get
```

Choose a backup and restore:

```bash
velero restore create --from-backup <backup-name>
```

> 💡 **Tip**: If you’re restoring to a new cluster with different node names or architecture, use Velero’s `--namespace-mappings` or other flags to handle discrepancies.

---

## 🛡️ Optional: Restore only certain namespaces or resources

```bash
velero restore create --from-backup <backup-name> \
  --include-namespaces my-namespace
```

---

## ✅ After Restoration

1. Check the status:

```bash
velero restore get
```

2. Validate critical resources:

```bash
kubectl get pods -A
kubectl get svc -A
kubectl get pvc -A
```

3. Ensure workloads are running and services are responding.

---

## 🧠 Final Thoughts

* **Velero is not a full etcd backup solution** – it's designed to restore Kubernetes resources by talking to the API server.
* Always combine **etcd snapshots** (taken via `etcdctl` or tools like `kubeadm`) **with Velero backups** for full disaster recovery coverage.
* Consider automating etcd snapshots and storing them alongside your Velero backups.

---

Would you like help creating a script to automate etcd + Velero backup and recovery together?
