Below is a **sample application Pod** that will be **correctly governed by this NetworkPolicy** and can **successfully use both the allowed ingress and egress rules**.

---

## 1. Namespace (required for ingress to work)

Your ingress rule requires traffic **from namespace labeled `name=dev`**.
So the DB pod itself should live in that namespace (typical setup).

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: dev
  labels:
    name: dev
```

---

## 2. Sample DB Application Pod (the pod being protected)

This pod **matches the NetworkPolicy’s `podSelector`** and can:

* Receive ingress on TCP 8080
* Send egress traffic to the allowed IP + port range

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: db-app
  namespace: dev
  labels:
    role: db-pod
spec:
  containers:
  - name: db-container
    image: nginx
    ports:
    - containerPort: 8080
```

### Why this pod satisfies the policy

| Policy Part                    | How this pod satisfies it                                 |
| ------------------------------ | --------------------------------------------------------- |
| `podSelector: role=db-pod`     | Pod label matches                                         |
| `policyTypes: Ingress, Egress` | Pod is now isolated                                       |
| Ingress port                   | Container listens on 8080                                 |
| Namespace selector             | Pod is in `dev` namespace                                 |
| Egress                         | Pod can initiate TCP connections to allowed IPs and ports |

---

## 3. Sample Client Pod (to test ingress access)

This pod is **not required**, but shows a pod that can **successfully connect to the DB pod**.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: internal-client
  namespace: dev
  labels:
    role: internal-db
spec:
  containers:
  - name: client
    image: curlimages/curl
    command: ["sleep", "3600"]
```

### Test ingress

```bash
kubectl exec -n dev internal-client -- curl http://db-app:8080
```

✅ Allowed by NetworkPolicy
❌ Same request from any other pod/namespace will be blocked

---

## 4. Egress example from the DB pod

From inside `db-app`, this is allowed:

```bash
curl 172.17.2.10:30080
```

This is blocked:

```bash
curl 172.17.1.10:30080   # excluded subnet
curl 8.8.8.8:53          # outside CIDR
```

---

## Summary

* **`db-app`** → satisfies **both ingress & egress rules**
* **`internal-client`** → valid ingress source
* Namespace labeling is **critical**
* Egress is tightly restricted to a CIDR + NodePort range

If you want, I can:

* Convert this to a **Deployment**
* Add **DNS egress support**
* Provide a **test matrix (allowed vs blocked traffic)**
