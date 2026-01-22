

**Amazon EKS does *not* install Metrics Server by default.**

What that means in practice:

* A fresh EKS cluster **does not have `metrics-server`**
* Commands like:

  ```bash
  kubectl top nodes
  kubectl top pods
  ```

  will **fail** until you install it
* HPA (Horizontal Pod Autoscaler) that relies on CPU/memory **won’t work** without it

### What *is* installed by default in EKS

EKS gives you the control plane and a few core addons, but **Metrics Server is not one of them**. Typical defaults include:

* kube-proxy
* CoreDNS
* VPC CNI

### How people usually install it

Most teams install Metrics Server via:

```bash
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```

Or as an **EKS add-on** (recommended for production):

```bash
aws eks create-addon \
  --cluster-name <cluster-name> \
  --addon-name metrics-server
```

### Quick check if it’s already installed

```bash
kubectl get deployment metrics-server -n kube-system
```

If you want, tell me:

* EKS version
* Whether you’re using managed node groups or Fargate

…and I’ll tell you the **best install method + flags** for your setup.


Great question — this is one of those “it just works… until it doesn’t” components 😄
Let’s break it down **clearly and EKS-specific**.

---

## What Metrics Server *does*

**Metrics Server collects short-term CPU and memory usage** for:

* **Nodes**
* **Pods / containers**

It exposes this data through the Kubernetes **Metrics API** (`metrics.k8s.io`).

That’s it.
👉 It **does NOT** store historical data
👉 It **does NOT** do alerting
👉 It **does NOT** replace Prometheus

---

## Why it exists (what uses it)

Metrics Server is mainly used by:

### 1️⃣ `kubectl top`

```bash
kubectl top nodes
kubectl top pods
```

Without Metrics Server → these commands fail.

---

### 2️⃣ Horizontal Pod Autoscaler (HPA)

HPA uses Metrics Server to decide:

* When to scale **up**
* When to scale **down**

Example:

```yaml
targetCPUUtilizationPercentage: 70
```

Metrics Server provides:

```
current CPU usage → HPA → scaling decision
```

---

## How Metrics Server works in **EKS**

Here’s the flow 👇

```
Kubelet (on each node)
   ↓
Metrics Server (kube-system)
   ↓
Metrics API (metrics.k8s.io)
   ↓
kubectl top / HPA
```

### Step-by-step

1. **Metrics Server runs as a Deployment** in `kube-system`
2. It **scrapes metrics from each node’s kubelet**

   * Endpoint: `https://<node-ip>:10250/metrics/resource`
3. Kubelet gets CPU/memory from:

   * cAdvisor
   * container runtime (containerd in EKS)
4. Metrics Server:

   * Aggregates the data
   * Keeps it in memory (≈ last 1–5 minutes)
5. Exposes it via the **Kubernetes Aggregation Layer**

   * API: `metrics.k8s.io`

---

## What’s EKS-specific about it

A few important EKS details:

### 🔐 TLS & auth

* Metrics Server authenticates to kubelet using **RBAC**
* In EKS, you often need:

  ```yaml
  --kubelet-insecure-tls
  ```

  because node certificates may not match IP/DNS perfectly

(Managed add-on handles this for you 👌)

---

### 🌐 Networking

* Metrics Server talks to nodes **inside the VPC**
* If using:

  * **Private clusters**
  * **Custom security groups**

  → Node SG must allow traffic on **10250**

---

### 🧠 Data limitations

Metrics Server:

* Only near-real-time data
* No disk, network, or custom metrics
* No long-term storage

That’s why people pair it with:

* **Prometheus**
* **CloudWatch Container Insights**

---

## What Metrics Server is *not*

Common misconception 🚫

| Thing              | Metrics Server         |
| ------------------ | ---------------------- |
| Monitoring         | ❌                      |
| Alerting           | ❌                      |
| Dashboards         | ❌                      |
| Historical metrics | ❌                      |
| Autoscaling        | ❌ (it only feeds data) |

It’s just a **metrics pipe**, not a monitoring system.

---

## Typical problems in EKS

If it breaks, you’ll usually see:

```bash
kubectl top nodes
# error: Metrics API not available
```

Common causes:

* Missing `--kubelet-insecure-tls`
* Security group blocking 10250
* Node not Ready
* Metrics Server pod crashing

---

## When you *need* it

You **must** install Metrics Server if you use:

* HPA (CPU / memory)
* `kubectl top`
* Some autoscaling tools (KEDA CPU/memory mode)

You **don’t need** it if:

* You only use Prometheus for observability
* No HPA based on CPU/memory

---

If you want, I can:

* Draw an **ASCII diagram** of the flow
* Show **real HPA → Metrics Server → Pod scaling**
* Compare **Metrics Server vs Prometheus vs CloudWatch**

Just say the word 👌
Got it — here’s a **real, end-to-end example** of
**HPA → Metrics Server → Pod scaling** in an **EKS cluster** 👇

---

## 1️⃣ The flow (big picture)

```
Pod CPU usage
   ↓
Kubelet (node)
   ↓
Metrics Server
   ↓
metrics.k8s.io API
   ↓
HPA controller
   ↓
Scale Deployment replicas
```

HPA checks metrics **every ~15 seconds** and decides whether to scale.

---

## 2️⃣ Example app (CPU-based scaling)

### Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cpu-demo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: cpu-demo
  template:
    metadata:
      labels:
        app: cpu-demo
    spec:
      containers:
      - name: app
        image: nginx
        resources:
          requests:
            cpu: "100m"
            memory: "128Mi"
          limits:
            cpu: "500m"
            memory: "256Mi"
```

⚠️ **Important**
HPA **requires `resources.requests.cpu`** — without it, scaling won’t work.

---

## 3️⃣ HPA definition (uses Metrics Server)

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: cpu-demo-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: cpu-demo
  minReplicas: 1
  maxReplicas: 5
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
```

### What this means

* Target: **70% CPU**
* CPU is calculated as:

  ```
  current CPU usage / requested CPU
  ```
* If average > 70% → scale up
* If average < 70% → scale down

---

## 4️⃣ What Metrics Server provides (real numbers)

Metrics Server reports something like:

```
Pod A: 90m CPU
Pod B: 80m CPU
```

Requests:

```
100m CPU per pod
```

HPA calculation:

```
Average utilization =
(90 + 80) / (2 × 100) = 85%
```

➡️ HPA decides to **scale up**

---

## 5️⃣ Scaling happens (real behavior)

```bash
kubectl get hpa
```

Output:

```
NAME           REFERENCE             TARGETS   MINPODS   MAXPODS   REPLICAS
cpu-demo-hpa   Deployment/cpu-demo   85%/70%   1         5         2
```

HPA increases replicas:

```
1 → 2 → 3 → ...
```

---

## 6️⃣ After load drops (scale down)

HPA:

* Waits for **stabilization window** (default ~5 min)
* Re-checks metrics
* Gradually scales down

This prevents **flapping**.

---

## 7️⃣ How often things happen

| Component             | Interval |
| --------------------- | -------- |
| Metrics Server scrape | ~15s     |
| HPA evaluation        | ~15s     |
| Scale down delay      | ~300s    |

---

## 8️⃣ Common EKS gotchas (very real)

### ❌ HPA shows `unknown`

```bash
kubectl describe hpa cpu-demo-hpa
```

Usually means:

* Metrics Server not installed
* Pod missing `resources.requests.cpu`
* Metrics Server can’t reach kubelet (10250 blocked)

---

### ❌ Scaling too aggressive

Fix with:

```yaml
behavior:
  scaleUp:
    stabilizationWindowSeconds: 60
  scaleDown:
    stabilizationWindowSeconds: 300
```

---

## 9️⃣ Summary (mental model)

* **Metrics Server = data source**
* **HPA = decision maker**
* **Deployment = thing that actually scales**

Metrics Server never scales anything by itself — it just reports numbers.

---

If you want next:

* 🔥 **Load-test demo** (`stress` / `hey`)
* 🧠 **Memory-based HPA**
* 🚀 **Custom metrics (Prometheus + HPA)**

Tell me what you want to dive into next.
