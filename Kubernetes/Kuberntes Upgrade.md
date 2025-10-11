Upgrading a Kubernetes cluster involves careful planning and execution to avoid downtime and ensure compatibility. The exact steps depend on **how your cluster is set up**, e.g.:

* **kubeadm**
* **managed services** (like GKE, EKS, AKS)
* **custom setups** (kops, Rancher, etc.)

I'll outline the general steps for **kubeadm**, which is the most commonly used tool for managing self-hosted Kubernetes clusters.

---

## 🔧 Prerequisites

* Take a **backup** of etcd and cluster configuration.
* Make sure **control plane and nodes are healthy** (`kubectl get nodes`, `kubectl get cs`).
* Know your current Kubernetes version (`kubectl version`).
* Upgrade in **minor version steps** (e.g., v1.27 → v1.28, not v1.27 → v1.29).
* Review the [Kubernetes Release Notes](https://kubernetes.io/docs/home/) for deprecations and changes.

---

## 📈 Step-by-Step Upgrade Using `kubeadm`

> Example: Upgrading from v1.27.x → v1.28.x

### 1. **Upgrade kubeadm on the control plane node**

```bash
sudo apt update
sudo apt install -y kubeadm=1.28.x-00
```

Replace `1.28.x-00` with the exact version.

Check version:

```bash
kubeadm version
```

### 2. **Plan the upgrade**

```bash
sudo kubeadm upgrade plan
```

This shows available versions and checks compatibility.

### 3. **Apply the upgrade**

```bash
sudo kubeadm upgrade apply v1.28.x
```

This upgrades the control plane components (API server, controller-manager, scheduler).

### 4. **Upgrade kubelet and kubectl**

```bash
sudo apt install -y kubelet=1.28.x-00 kubectl=1.28.x-00
sudo systemctl restart kubelet
```

---

## 👷 Upgrade Worker Nodes

For **each worker node**, do the following:

### 1. Drain the node

```bash
kubectl drain <node-name> --ignore-daemonsets --delete-emptydir-data
```

### 2. Upgrade `kubeadm`

```bash
sudo apt install -y kubeadm=1.28.x-00
sudo kubeadm upgrade node
```

### 3. Upgrade `kubelet` and `kubectl`

```bash
sudo apt install -y kubelet=1.28.x-00 kubectl=1.28.x-00
sudo systemctl restart kubelet
```

### 4. Uncordon the node

```bash
kubectl uncordon <node-name>
```

Repeat for all worker nodes.

---

## ✅ Post-upgrade checklist

* Verify all nodes are `Ready`:

  ```bash
  kubectl get nodes
  ```

* Check core components:

  ```bash
  kubectl get pods -n kube-system
  ```

* Run functional tests (if you have a test suite).

---

## 🛠️ For Managed Clusters

If you're using a managed Kubernetes service:

* **GKE**: Use the Google Cloud Console or `gcloud container clusters upgrade`
* **EKS**: Use the AWS Console or `eksctl upgrade cluster`
* **AKS**: Use Azure Portal or `az aks upgrade`

---

If you tell me which platform or tool you’re using (kubeadm, kops, EKS, etc.), I can tailor the instructions precisely.
