markdown
# AWS EBS Configuration for Kubernetes (2026)

This guide provides the standard YAML manifests for using [Amazon EBS CSI driver](https://github.com) to provide persistent storage for your Kubernetes workloads.

## 1. StorageClass
The `StorageClass` defines the EBS volume type (gp3 is the modern standard) and ensures the volume is provisioned in the same Availability Zone as the Pod.

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: ebs-sc
provisioner: ebs.csi.aws.com
parameters:
  type: gp3
  fsType: ext4
volumeBindingMode: WaitForFirstConsumer
allowVolumeExpansion: true
Use code with caution.

2. PersistentVolumeClaim (PVC)
The PVC requests a specific size of storage from the StorageClass defined above.
yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: ebs-claim
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: ebs-sc
  resources:
    requests:
      storage: 10Gi
Use code with caution.

3. Pod Deployment
The Pod references the ebs-claim to mount the EBS volume into the container's file system.
yaml
apiVersion: v1
kind: Pod
metadata:
  name: app-server
spec:
  containers:
  - name: nginx
    image: nginx
    volumeMounts:
    - name: storage-volume
      mountPath: /var/www/html
  volumes:
  - name: storage-volume
    persistentVolumeClaim:
      claimName: ebs-claim
Use code with caution.

Implementation Notes
Prerequisite: Ensure the Amazon EBS CSI Driver is installed in your cluster.
Availability Zones: Using volumeBindingMode: WaitForFirstConsumer is critical. It prevents the "volume in different AZ" error by waiting until the Pod is scheduled before creating the volume.
Access Modes: EBS only supports ReadWriteOnce (RWO). If you require multi-node access (RWX), consider using Amazon EFS.
Expansion: To resize a volume in 2026, simply update the storage request in the PVC and re-apply. The CSI driver handles the online filesystem expansion.
{content: }
