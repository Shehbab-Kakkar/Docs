For the **same environment variable name** in a Kubernetes Pod/Deployment, the precedence is **clear and strict**:

### ✅ Precedence (highest → lowest)

1. **`env`**
2. **`envFrom`**

So:

* If a variable is defined in **both `env` and `envFrom`**, the value in **`env` wins**.
* Kubernetes **does not merge** values — it simply overrides.

### Example

```yaml
envFrom:
  - configMapRef:
      name: app-config
env:
  - name: DB_HOST
    value: "db.prod.svc.cluster.local"
```

ConfigMap (`app-config`):

```yaml
DB_HOST=db.dev.svc.cluster.local
```

👉 **Final value inside the container**

```text
DB_HOST=db.prod.svc.cluster.local
```

### Important notes

* If the **same variable name appears in multiple `envFrom` sources**, the **last one listed wins**
* If the same variable appears **multiple times in `env`**, the **last definition wins**
* Kubernetes does **not raise an error** for duplicates

### Summary

| Source             | Same variable name |
| ------------------ | ------------------ |
| `env` vs `envFrom` | **`env` wins**     |
| Multiple `envFrom` | **Last one wins**  |
| Multiple `env`     | **Last one wins**  |

If you want, I can show the exact behavior from the Kubernetes spec or a quick `kubectl exec env` demo.
