Here's how each term applies during **deployment**:

---

**progressDeadlineSeconds**  
- **Explanation:**  
  During deployment, Kubernetes uses `progressDeadlineSeconds` to determine how long it should wait for the deployment to become successful before reporting a failure.  
- **Example:**  
  If you set `progressDeadlineSeconds: 300`, Kubernetes considers the deployment failed if it hasn't finished rolling out (i.e., reached the desired state) after 300 seconds (5 minutes).  
- **Usage:**  
  Helps ensure you are notified of slow or stuck deployments so corrective actions can be taken.

---

**revisionHistoryLimit**  
- **Explanation:**  
  This setting controls how many old ReplicaSet objects (previous deployment versions) Kubernetes keeps for a particular deployment.  
- **Example:**  
  If you set `revisionHistoryLimit: 3`, only the last 3 ReplicaSets (old revisions) will be retained; earlier ones are deleted to free resources.  
- **Usage:**  
  Useful for rollbacks: you can only roll back to the number of old revisions specified. Helps conserve cluster resources and avoid clutter.

---

### In Context

When you create or update a deployment manifest in Kubernetes:
```yaml
spec:
  progressDeadlineSeconds: 300
  revisionHistoryLimit: 3
```
- This **ensures deployments are monitored**, and you keep a manageable number of rollback targets.

---

**Summary Table:**

| Setting                 | Deployment Role                                                                                 |
|-------------------------|------------------------------------------------------------------------------------------------|
| progressDeadlineSeconds | Time to wait for deployment success before marking as failed                                   |
| revisionHistoryLimit    | Number of previous deployment versions kept for rollbacks and history                          |
