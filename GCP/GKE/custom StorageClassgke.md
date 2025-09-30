To create a custom StorageClass in Google Kubernetes Engine (GKE), follow these steps:

## 1. Choose a Provisioner

GKE supports several provisioners, but the most common is `kubernetes.io/gce-pd` (for standard persistent disks) or `pd.csi.storage.gke.io` (for the CSI driver, which is recommended).

## 2. Create a StorageClass YAML Manifest

Here’s an example using the **GKE CSI driver** (`pd.csi.storage.gke.io`), which you should use for new clusters:

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: custom-gke-storage
provisioner: pd.csi.storage.gke.io
parameters:
  type: pd-ssd # or pd-standard
  # You can add more parameters here, see docs below
reclaimPolicy: Delete
volumeBindingMode: WaitForFirstConsumer
allowVolumeExpansion: true
```

**Parameters:**
- `type`: Storage type (`pd-standard` for HDD, `pd-ssd` for SSD)
- `reclaimPolicy`: `Delete` (default) or `Retain`
- `volumeBindingMode`: Usually `WaitForFirstConsumer` for better scheduling

## 3. Apply the StorageClass

Save the above YAML to a file, e.g., `custom-storageclass.yaml`, then run:

```sh
kubectl apply -f custom-storageclass.yaml
```

## 4. Use the StorageClass in a PersistentVolumeClaim

Reference your StorageClass in a PVC:

```yaml
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
  storageClassName: custom-gke-storage
```

## 5. Apply the PVC

```sh
kubectl apply -f my-pvc.yaml
```

---

## References

- [GKE CSI StorageClass documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/gce-pd-csi-driver#create_storage_class)
- [Kubernetes StorageClass docs](https://kubernetes.io/docs/concepts/storage/storage-classes/)

Let me know if you need a specific StorageClass example or have a particular use-case (e.g., regional disks, encryption, etc)!
