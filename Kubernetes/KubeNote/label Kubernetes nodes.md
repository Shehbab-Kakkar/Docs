To label Kubernetes nodes, you use the **`kubectl label`** command. Here’s a clear, practical guide.

---

## 1. List your nodes

First, see the node names:

```bash
kubectl get nodes
```

Example output:

```
NAME            STATUS   ROLES    AGE   VERSION
worker-node-1   Ready    <none>   10d   v1.29.0
worker-node-2   Ready    <none>   10d   v1.29.0
```

---

## 2. Add a label to a node

Basic syntax:

```bash
kubectl label nodes <node-name> <key>=<value>
```

Example:

```bash
kubectl label nodes worker-node-1 disktype=ssd
```

✅ Output:

```
node/worker-node-1 labeled
```

---

## 3. Verify the label

You can check labels in several ways.

### Show all labels for a node

```bash
kubectl describe node worker-node-1
```

### Show labels in a compact view

```bash
kubectl get nodes --show-labels
```

---

## 4. Overwrite an existing label

If the label key already exists, Kubernetes will block it unless you allow overwrite:

```bash
kubectl label nodes worker-node-1 disktype=hdd --overwrite
```

---

## 5. Remove a label

Use a dash (`-`) after the key:

```bash
kubectl label nodes worker-node-1 disktype-
```

---

## 6. Use labels for scheduling (example)

Once labeled, you can schedule Pods onto specific nodes.

### Pod using `nodeSelector`

```yaml
spec:
  nodeSelector:
    disktype: ssd
```

### (Better) using `nodeAffinity`

```yaml
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: disktype
            operator: In
            values:
            - ssd
```

---

## 7. Common label naming conventions

* Use lowercase
* Use DNS-style prefixes for custom labels:

  ```text
  example.com/environment=prod
  ```
* Avoid changing built-in labels like `kubernetes.io/hostname`

---

If you want, tell me:

* Your Kubernetes distro (EKS, GKE, k3s, on-prem, etc.)
* What you want to use the labels for (GPU nodes, zones, workloads)

I can tailor examples exactly to your setup.
