**kubelet** is the **node-level agent in Kubernetes**.
It runs on **every worker node** and is responsible for making sure the containers on that node are running exactly as Kubernetes expects.

---

## What kubelet does (simple explanation)

> kubelet watches the Kubernetes API and **turns Pod specs into running containers** on a node.

---

## Core Responsibilities

### 1. Pod Lifecycle Management

* Watches the API server for Pods scheduled to its node
* Pulls container images
* Starts, stops, and restarts containers
* Ensures the Pod matches the desired state

---

### 2. Communicates with Container Runtime

* Talks to container runtimes via **CRI (Container Runtime Interface)**
* Examples:

  * containerd
  * CRI-O

kubelet **does not run containers itself** — it tells the runtime what to do.

---

### 3. Health Checks

* Executes **liveness, readiness, and startup probes**
* Restarts containers if probes fail
* Reports container health to the API server

---

### 4. Node Status Reporting

* Reports:

  * CPU, memory, disk usage
  * Node conditions (Ready, DiskPressure, etc.)
* Sends heartbeats to the control plane

---

### 5. Volume Management

* Attaches, mounts, and unmounts volumes
* Works with CSI drivers
* Ensures volumes are available before containers start

---

### 6. Resource Management

* Enforces CPU and memory requests/limits
* Works with cgroups
* Handles OOM events

---

### 7. Security & Secrets

* Fetches Secrets and ConfigMaps
* Sets up credentials and certificates
* Applies security contexts

---

## What kubelet does NOT do

❌ It does **not** schedule Pods
❌ It does **not** expose Services
❌ It does **not** manage networking (handled by CNI)
❌ It does **not** store cluster state

---

## kubelet in the Pod lifecycle

```text
User → kubectl apply
        ↓
API Server
        ↓
Scheduler (chooses node)
        ↓
kubelet (on that node)
        ↓
Container Runtime → Containers running
```

---

## Key kubelet Interfaces

| Interface  | Purpose                 |
| ---------- | ----------------------- |
| CRI        | Container runtime       |
| CNI        | Networking (indirectly) |
| CSI        | Storage                 |
| API Server | Cluster state           |

---

## One-line summary

> **kubelet is the “node manager” that ensures Pods are running correctly on each Kubernetes node.**

If you want, I can explain:

* kubelet vs kube-proxy
* kubelet vs scheduler
* kubelet interview questions
* kubelet troubleshooting scenarios
