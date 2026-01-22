AKS sits **between EKS and GKE** — more managed than EKS, less “invisible” than GKE.

---

## Is Metrics Server installed by default in AKS?

**Yes, in most modern AKS clusters.**

* AKS **deploys Metrics Server automatically**
* `kubectl top nodes/pods` works out of the box
* HPA (CPU & memory) works without manual install

👉 You usually **don’t install it yourself** anymore.

---

## How HPA → Metrics works in AKS

Same Kubernetes pattern again:

```
Kubelet
  ↓
Metrics Server (AKS-managed)
  ↓
metrics.k8s.io
  ↓
HPA
  ↓
Scale workload
```

---

## What’s different vs EKS and GKE

### 1️⃣ Visible but managed

Unlike GKE:

* You **can see** `metrics-server` in `kube-system`
* But it’s **managed by AKS**
* You should **not modify or upgrade it manually**

```bash
kubectl get deployment metrics-server -n kube-system
```

---

### 2️⃣ Secure kubelet access (no hacks)

AKS configures:

* TLS correctly
* RBAC correctly
* Kubelet ports correctly

So:

* ❌ No `--kubelet-insecure-tls`
* ❌ No port 10250 drama

This already saves hours compared to EKS 😅

---

### 3️⃣ Works automatically with HPA

This just works:

```yaml
metrics:
- type: Resource
  resource:
    name: cpu
    target:
      type: Utilization
      averageUtilization: 70
```

No extra steps.

---

## AKS + Azure Monitor (important distinction)

Same story as other clouds:

| Feature            | Metrics Server | Azure Monitor |
| ------------------ | -------------- | ------------- |
| HPA CPU/memory     | ✅              | ❌             |
| `kubectl top`      | ✅              | ❌             |
| Dashboards         | ❌              | ✅             |
| Alerts             | ❌              | ✅             |
| Historical metrics | ❌              | ✅             |

Azure Monitor is **not** used by HPA unless you wire up **custom metrics**.

---

## AKS quirks & gotchas

### ❗ Requests are mandatory

Just like everywhere else:

```yaml
resources:
  requests:
    cpu: "100m"
```

No request → HPA won’t scale.

---

### ❗ Older clusters

Very old AKS versions:

* Metrics Server might not be enabled
* Needed manual install

Check:

```bash
kubectl top nodes
```

If it fails → Metrics Server isn’t running.

---

## Quick comparison: EKS vs GKE vs AKS

| Feature                | EKS   | AKS | GKE |
| ---------------------- | ----- | --- | --- |
| Metrics Server default | ❌     | ✅   | ✅   |
| Managed by provider    | ❌     | ✅   | ✅   |
| Visible to user        | ✅     | ✅   | ❌   |
| Needs insecure TLS     | Often | ❌   | ❌   |
| HPA out-of-box         | ❌     | ✅   | ✅   |

---

## TL;DR

* **EKS** → you install & manage Metrics Server
* **AKS** → installed and managed, visible
* **GKE** → fully hidden and managed
* **HPA behavior is identical** across all three

If you want, next we can:

* Compare **autoscaling maturity** across clouds
* Show **custom metrics HPA** (Azure Monitor vs CloudWatch vs GCM)
* Do a **cloud-agnostic HPA best-practice checklist**

Just say the word 🚀
