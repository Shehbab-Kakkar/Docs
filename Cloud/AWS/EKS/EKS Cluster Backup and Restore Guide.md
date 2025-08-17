# 🛡️ EKS Cluster Backup and Restore Guide

This guide covers best practices and examples for backing up and restoring your Amazon EKS (Elastic Kubernetes Service) cluster, including both Kubernetes resources (YAML objects) and persistent data (volumes).

---

## 🏗️ What to Back Up in EKS?

1. **Kubernetes Resources:**  
   - Deployments, Services, ConfigMaps, Secrets, Ingresses, CRDs, etc.
   - Typically stored in the Kubernetes API server (etcd).

2. **Persistent Volumes (Data):**  
   - EBS volumes, EFS, or S3 buckets mounted to pods.
   - Application data (databases, user uploads, etc.)

---

## 🚀 Recommended Tools

- **[Velero](https://velero.io/):** Open-source tool for backup/restore of Kubernetes resources and persistent volumes.
- **EBS Snapshots:** For persistent data in EBS volumes.
- **kubectl:** For simple YAML export/import (not suitable for large/production setups).

---

## ✅ Backup & Restore with Velero (Recommended)

### 1️⃣ Install Velero (CLI & Server)

#### a. Install Velero CLI

```bash
curl -Lo velero.tar.gz https://github.com/vmware-tanzu/velero/releases/download/v1.13.1/velero-v1.13.1-linux-amd64.tar.gz
tar -xvf velero.tar.gz
sudo mv velero-v1.13.1-linux-amd64/velero /usr/local/bin/
velero version
```

#### b. Install Velero on EKS

```bash
velero install \
  --provider aws \
  --plugins velero/velero-plugin-for-aws:v1.9.2 \
  --bucket <your-backup-bucket> \
  --backup-location-config region=<region> \
  --snapshot-location-config region=<region> \
  --secret-file ./credentials-velero
```

- `<your-backup-bucket>`: S3 bucket for storing backups.
- `./credentials-velero`: AWS credentials file for Velero.

---

### 2️⃣ Take a Backup

Backup all namespaces:

```bash
velero backup create my-cluster-backup --include-namespaces '*'
```

Backup a specific namespace:

```bash
velero backup create my-namespace-backup --include-namespaces my-namespace
```

Include persistent volumes:

- Velero by default can snapshot EBS volumes if permissions are set.

Check status:

```bash
velero backup get
```

---

### 3️⃣ Restore from Backup

Restore everything:

```bash
velero restore create --from-backup my-cluster-backup
```

Restore a specific namespace:

```bash
velero restore create --from-backup my-namespace-backup --include-namespaces my-namespace
```

Check restore status:

```bash
velero restore get
```

---

## 🛠️ Manual Backup/Restore (Small/Test Clusters)

### Export all resources in all namespaces:

```bash
kubectl get all --all-namespaces -o yaml > all-resources.yaml
```

### Restore:

```bash
kubectl apply -f all-resources.yaml
```

> **Note:** This does NOT back up persistent data in volumes.

---

## 💾 Persistent Data: EBS/EFS Snapshots

- For workloads using EBS:
  - [Take EBS snapshots](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-creating-snapshot.html) (automate with AWS Backup or Lambda).
  - Restore by creating a new volume from the snapshot and re-attaching to pods.

- For EFS:
  - Use EFS backup (AWS Backup or EFS native).

---

## 🔄 Disaster Recovery Checklist

- Regularly **test restores** in a separate cluster.
- Backup **both** Kubernetes objects and persistent data.
- Store backup artifacts (YAML, Velero, EBS/EFS snapshots) in another region/account for extra safety.
- Use versioned S3 buckets for backup storage.

---

## 📚 References

- [Velero Documentation](https://velero.io/docs/)
- [EKS Backup Best Practices (AWS)](https://aws.amazon.com/blogs/containers/backup-and-restore-methods-for-amazon-eks/)
- [AWS EBS Snapshots](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-creating-snapshot.html)
- [AWS EFS Backups](https://docs.aws.amazon.com/efs/latest/ug/awsbackup.html)

---

**Tip:**  
Automate your backup jobs and monitor Velero/EBS snapshot status for reliability!
