This is one of those **“there’s no single right number”** questions — but big enterprises *do* converge on very similar **patterns**. I’ll give you **realistic, production-grade examples** and explain the **why**, not just numbers.

I’ll assume:

* Large enterprise
* Multiple teams
* Shared clusters
* Strong guardrails
* Separate namespaces per app/team
* `prod` is mission-critical

---

## First: how big enterprises think about quotas

They usually layer **3 levels of control**:

1. **Cluster capacity** (e.g. 500–2000 vCPU, TBs of RAM)
2. **Namespace quotas** (hard limits)
3. **Pod requests/limits** (guaranteed scheduling + burst control)

Quotas are **guardrails**, not exact sizing.

---

## Typical enterprise PROD namespace quotas

### 🟥 Tier 1 – Core production apps (revenue / customer-facing)

**Per namespace (example):**

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: prod-quota
  namespace: payments-prod
spec:
  hard:
    requests.cpu: "200"
    requests.memory: "400Gi"
    limits.cpu: "400"
    limits.memory: "800Gi"
    pods: "500"
```

**What this means**

* Guaranteed capacity: **200 vCPU / 400Gi**
* Max burst: **400 vCPU / 800Gi**
* Enough headroom for spikes
* Prevents one team from nuking the cluster

Used by:

* Payments
* Checkout
* Auth
* Core APIs

---

### 🟧 Tier 2 – Important but non-core prod apps

```yaml
requests.cpu: "50"
requests.memory: "100Gi"
limits.cpu: "100"
limits.memory: "200Gi"
pods: "200"
```

Used by:

* Internal APIs
* Reporting services
* Admin backends

---

### 🟨 Tier 3 – Shared / support prod workloads

```yaml
requests.cpu: "10"
requests.memory: "20Gi"
limits.cpu: "20"
limits.memory: "40Gi"
pods: "100"
```

Used by:

* CronJobs
* Workers
* Low-traffic services

---

## Namespace-level LimitRange (almost always enforced)

Big enterprises **always** pair quotas with `LimitRange`:

```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: prod-limits
  namespace: payments-prod
spec:
  limits:
  - type: Container
    defaultRequest:
      cpu: "250m"
      memory: "512Mi"
    default:
      cpu: "1"
      memory: "2Gi"
    max:
      cpu: "8"
      memory: "16Gi"
```

Why:

* Prevents “no request set” pods
* Prevents insane container sizes
* Forces predictable scheduling

---

## CPU & Memory ratios enterprises expect

Very common rules of thumb:

| Metric                 | Typical ratio         |
| ---------------------- | --------------------- |
| CPU limit : request    | 2:1                   |
| Memory limit : request | 2:1 (sometimes 1.5:1) |
| CPU per pod            | 250m–2000m            |
| Memory per pod         | 512Mi–4Gi             |

Memory overcommit is **dangerous**, so enterprises are stricter there.

---

## Cluster sizing example (realistic)

**Prod cluster**

* 200 nodes
* 8 vCPU / 32Gi each

Total:

* 1600 vCPU
* 6.4 TiB RAM

Enterprise might allocate:

* 70% for prod namespaces
* 20% for system / platform
* 10% buffer

So:

* ~1100 vCPU allocatable to prod namespaces
* Quotas enforce fair sharing

---

## What enterprises NEVER do

❌ Unlimited namespaces
❌ No ResourceQuota in prod
❌ No requests/limits
❌ One giant prod namespace
❌ Same quotas for dev/stage/prod

---

## Environment comparison (typical)

| Env   | CPU quota    | Memory quota |
| ----- | ------------ | ------------ |
| dev   | 5–10 vCPU    | 10–20Gi      |
| stage | 20–50 vCPU   | 50–100Gi     |
| prod  | 50–200+ vCPU | 100–400+Gi   |

---

## TL;DR (executive summary)

* **Prod namespaces usually get 50–200 vCPU**
* Memory scales roughly **2–4 Gi per vCPU**
* CPU limit:request ≈ **2:1**
* Every prod namespace has:

  * `ResourceQuota`
  * `LimitRange`
* Numbers scale by **business criticality**, not team size

---

If you want, I can:

* Design **quotas for your exact cluster size**
* Give **Helm-friendly quota templates**
* Show **platform vs app namespace separation**
* Map quotas to **SRE error budgets**

Just tell me:

* Cluster size
* Cloud provider
* How many prod namespaces / teams
