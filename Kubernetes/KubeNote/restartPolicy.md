There is **no single “best” restart policy** in Kubernetes — the *best policy depends entirely on the workload type*.

Here’s the **correct choice by use case**, which is how Kubernetes is designed to be used.

---

## Best restart policy by workload type

### ✅ **Long-running services (most common)**

**Best policy:** `Always`

**Use for:**

* Web servers
* APIs
* Microservices
* Background workers
* Deployments / ReplicaSets / StatefulSets

**Why:**

* Services must stay running
* Containers are automatically restarted on crash
* Required by Deployments

```yaml
restartPolicy: Always
```

> ⭐ **This is the default and most commonly “best” policy**

---

### ✅ **Batch jobs / one-time tasks**

**Best policy:** `OnFailure`

**Use for:**

* Jobs
* CronJobs
* Data processing
* ETL
* Backups
* Migrations

**Why:**

* Retry only when something goes wrong
* Don’t rerun successful tasks

```yaml
restartPolicy: OnFailure
```

---

### ✅ **Run once, no retries**

**Best policy:** `Never`

**Use for:**

* Debug pods
* Manual testing
* One-time scripts
* Experiments

**Why:**

* You want the exact result of a single execution
* No automatic retries

```yaml
restartPolicy: Never
```

---

## Quick decision table

| Workload         | Best restartPolicy |
| ---------------- | ------------------ |
| Web app / API    | `Always`           |
| Microservice     | `Always`           |
| Batch Job        | `OnFailure`        |
| CronJob          | `OnFailure`        |
| Migration        | `OnFailure`        |
| Debug / Test Pod | `Never`            |

---

## Kubernetes best practice (important)

> ❌ **Don’t use bare Pods for production**
>
> ✅ Use **controllers**:

* Deployment → `Always`
* Job / CronJob → `OnFailure`
* StatefulSet → `Always`

The controller, not the Pod, handles reliability.

---

## Final answer (short)

* **Best overall (most workloads):** `Always`
* **Best for batch work:** `OnFailure`
* **Best for debugging:** `Never`

