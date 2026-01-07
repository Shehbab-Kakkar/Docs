Here’s an example of **rolling back a deployment** using `revisionHistoryLimit` in Kubernetes:

---

### 1. **Deployment Manifest with revisionHistoryLimit**

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app
spec:
  revisionHistoryLimit: 3
  replicas: 2
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: my-app
        image: my-app:v1
```

---

### 2. **Update the Deployment (New Revision)**

Suppose you update the image version:
```yaml
      containers:
      - name: my-app
        image: my-app:v2
```
Kubernetes creates a new ReplicaSet for this revision. Old ReplicaSets are kept, up to the limit set by `revisionHistoryLimit`.

---

### 3. **Rolling Back the Deployment**

To rollback to the previous revision, use the following command:
```bash
kubectl rollout undo deployment/my-app
```
Kubernetes will use one of the previous revisions (as stored in the ReplicaSets, up to the revisionHistoryLimit) to restore the deployment state.

---

### 4. **Impact of revisionHistoryLimit**

- If you have `revisionHistoryLimit: 3`, Kubernetes only keeps the last 3 revisions.
- If you update your deployment multiple times (let’s say 5 times), only the last 3 will be available for rollback; older ones are deleted.

---

### **Summary Table:**

| Action                 | Outcome                         |
|------------------------|---------------------------------|
| Deploy v1              | Revision 1 stored               |
| Update to v2           | Revision 2 stored               |
| Update to v3           | Revision 3 stored               |
| Update to v4           | Revision 4 stored, revision 1 deleted (if revisionHistoryLimit: 3) |
| Rollback command       | `kubectl rollout undo deployment/my-app` brings deployment back to v3 |

---

To **check all revisions** and **rollback to a specific revision (e.g., revision 2)** for a Kubernetes Deployment, follow these steps:

---

### 1. **Check Deployment Revisions**

List all ReplicaSets for your deployment (each represents a revision):

```bash
kubectl get rs -l app=my-app
```
or (shows more details):

```bash
kubectl get rs --selector=app=my-app --output=wide
```

Alternatively, to see the deployment history via rollout:

```bash
kubectl rollout history deployment/my-app
```

You’ll get output like:
```
REVISION  CHANGE-CAUSE
1         <none>
2         <none>
3         <none>
```

---

### 2. **Rollback to a Specific Revision (e.g., revision 2)**

Use the following command:
```bash
kubectl rollout undo deployment/my-app --to-revision=2
```

This will change your deployment to revision 2, if it is still retained according to `revisionHistoryLimit`.

---

### **Notes**
- The ability to rollback to a specific revision depends on whether it is still present (i.e., not deleted due to `revisionHistoryLimit`).
- To see details of a specific revision:
  ```bash
  kubectl rollout history deployment/my-app --revision=2
  ```

---

**Summary Table:**

| Step                                  | Command                                                         |
|----------------------------------------|-----------------------------------------------------------------|
| List all revisions                     | `kubectl rollout history deployment/my-app`                     |
| See details for a specific revision    | `kubectl rollout history deployment/my-app --revision=2`        |
| Rollback to revision 2                 | `kubectl rollout undo deployment/my-app --to-revision=2`        |

**Tip:** Always make sure the revision you want to rollback to is available (`revisionHistoryLimit`). Otherwise, rollback won’t work for old revisions.
**In short:**  
`revisionHistoryLimit` controls how far back you can rollback.  
If the revision you want to revert to has been deleted (because it’s older than the limit), you cannot rollback to it.
