A **GKE (Google Kubernetes Engine) cluster is *regional*, not global**—but it helps to separate *control plane*, *nodes*, and *service exposure*, because that’s where confusion usually comes from.

---

## How GKE scope works

### 1. **Zonal clusters**

* **Scoped to a single zone** (e.g. `us-central1-a`)
* Control plane runs in that zone
* Worker nodes run in that zone
* Less highly available
* ❌ Not global

---

### 2. **Regional clusters (recommended)**

* **Scoped to a single region** (e.g. `us-central1`)
* Control plane is **replicated across multiple zones** in the region
* Worker nodes run in multiple zones within that region
* Higher availability than zonal clusters
* ❌ Still *not global* (cannot span multiple regions)

---

## What *can* be global with GKE

While the **cluster itself is regional**, some related components can be **global**:

### ✅ Global load balancing

* GKE integrates with **Google Cloud HTTP(S) Load Balancer**
* This load balancer is **global**
* Traffic can be routed to:

  * Multiple services
  * Multiple GKE clusters
  * Even multiple regions (via multi-cluster setups)

### ✅ Multi-cluster architectures

You can build *global* apps by combining multiple regional clusters:

* Multiple GKE clusters in different regions
* Global HTTP(S) Load Balancer in front
* Optional:

  * **GKE Multi-Cluster Ingress**
  * **Fleet / Anthos**
  * **Traffic Director**

This gives you **global availability**, but the clusters themselves remain regional.

---

## Quick summary

| Component                        | Scope                 |
| -------------------------------- | --------------------- |
| GKE cluster                      | **Regional or zonal** |
| Control plane (regional cluster) | Regional (multi-zone) |
| Worker nodes                     | Regional (multi-zone) |
| GKE service (ClusterIP/NodePort) | Cluster-local         |
| GKE Ingress / HTTP(S) LB         | **Global**            |

---

### One-line answer

> **A GKE cluster is regional (or zonal), not global—but you can build globally available services on top of it.**

If you want, I can help you design:

* a **multi-region GKE architecture**
* or explain **when to choose zonal vs regional** clusters
