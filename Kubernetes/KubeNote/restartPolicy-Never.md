In Kubernetes, a **Pod’s `restartPolicy`** controls what the kubelet should do when a **container inside the Pod exits**.

---

## `restartPolicy: Never` — Explanation

When a Pod is created with:

```yaml
restartPolicy: Never
```

### What it means

* **Containers are NOT restarted** if they exit (whether they succeed or fail).
* Once a container stops, it stays stopped.
* The Pod remains in a **terminal phase**.

---

## Pod lifecycle behavior

| Container Exit Reason     | Pod Status  |
| ------------------------- | ----------- |
| Exit code `0` (success)   | `Succeeded` |
| Exit code ≠ `0` (failure) | `Failed`    |

No restart attempts are made in either case.

---

## Common use cases

`restartPolicy: Never` is typically used for:

* **Batch jobs**
* **One-time tasks**
* **Data migrations**
* **Debug / test pods**
* **CI/CD steps**

Example scenarios:

* Run a script once and stop
* Execute a database migration
* Perform a backup job

> ⚠️ Note: In practice, Kubernetes **Jobs** are preferred over bare Pods for these cases.

---

## Example Pod YAML

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: run-once-pod
spec:
  restartPolicy: Never
  containers:
  - name: app
    image: busybox
    command: ["sh", "-c", "echo Hello Kubernetes && exit 0"]
```

After execution:

* Pod status → `Succeeded`
* Container → Not restarted

---

## Comparison with other restart policies

| Policy             | Behavior                                        |
| ------------------ | ----------------------------------------------- |
| `Always` (default) | Always restart containers (used by Deployments) |
| `OnFailure`        | Restart only if exit code ≠ 0                   |
| `Never`            | Never restart containers                        |

---

## Important notes

* `restartPolicy` is **Pod-level**, not container-level.
* Controllers enforce rules:

  * **Deployment / ReplicaSet** → must use `Always`
  * **Job** → allows `OnFailure` or `Never`
* Even with `Never`, a **new Pod may be created** by a controller (like a Job), but the **same Pod is never restarted**.

---

## Summary

`restartPolicy: Never` means:

> *“Run the container once. If it stops, do nothing.”*

It’s best suited for **run-once workloads**, especially when used with a **Job** rather than a standalone Pod.

If you want, I can also explain:

* Difference between `Never` and `OnFailure`
* How Jobs behave with `restartPolicy`
* Real-world Job YAML examples
