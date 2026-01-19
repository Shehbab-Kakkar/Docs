In Kubernetes, the **master node** (now commonly called the **control plane**) runs the components that **manage and control the cluster**.

---

## Core Components Running on the Kubernetes Master (Control Plane)

### 1. **kube-apiserver**

* **Front door of the cluster**
* Exposes the Kubernetes API (REST)
* Handles authentication, authorization, and admission control
* All components communicate **only through the API server**

---

### 2. **etcd**

* Distributed **key-value store**
* Stores the **entire cluster state**:

  * Pods, Services, ConfigMaps, Secrets
  * Node info, RBAC, CRDs
* Strongly consistent and highly available

---

### 3. **kube-scheduler**

* Decides **which node a Pod should run on**
* Considers:

  * Resource availability
  * Node affinity / taints & tolerations
  * Pod affinity / anti-affinity
* Binds Pods to nodes via the API server

---

### 4. **kube-controller-manager**

* Runs multiple controllers as a single process
* Continuously reconciles desired vs actual state

Controllers include:

* Node Controller
* Replication Controller
* Deployment Controller
* Endpoint Controller
* Namespace Controller
* ServiceAccount Controller

---

### 5. **cloud-controller-manager** (optional)

* Integrates Kubernetes with **cloud provider APIs**
* Handles cloud-specific resources:

  * Load balancers
  * Volumes
  * Node lifecycle (cloud instances)

Used in cloud environments (AWS, GCP, Azure).

---

## Control Plane Component Summary

| Component                | Responsibility        |
| ------------------------ | --------------------- |
| kube-apiserver           | API entry point       |
| etcd                     | Cluster state storage |
| kube-scheduler           | Pod placement         |
| kube-controller-manager  | State reconciliation  |
| cloud-controller-manager | Cloud integration     |

---

## Important Notes

* Control plane components usually run as **static Pods**
* They run in the `kube-system` namespace
* In HA setups, components are replicated across multiple masters
* Worker nodes do **not** run these components

---

## Simple Flow

```text
kubectl → kube-apiserver → etcd
                    ↓
        scheduler / controllers
                    ↓
              kubelet (worker nodes)
```

---

## One-line summary

> **The Kubernetes master (control plane) runs the API server, etcd, scheduler, and controllers that manage the entire cluster.**

