# 🔍 Checking for Istio CNI DaemonSet on All Nodes in GKE (`kube-system` namespace)

To verify that **Istio CNI** is deployed as a DaemonSet and running on all three nodes in your GKE cluster (specifically in the `kube-system` namespace), use the following steps:

---

## 1. **Get DaemonSets in the kube-system Namespace**

```bash
kubectl get daemonset.apps -n kube-system
```

## 2. **Filter for Istio CNI**

To narrow it down to Istio CNI:

```bash
kubectl get daemonset.apps -n kube-system | grep -i istio
```

*Typical output if Istio CNI is deployed:*

```text
istio-cni-node   3   3   3   3   0   10m
```

- The columns (DESIRED, CURRENT, READY, UP-TO-DATE, AVAILABLE, AGE) should all show `3` on a 3-node cluster if healthy.

---

## 3. **Describe the DaemonSet for More Details**

```bash
kubectl -n kube-system describe daemonset istio-cni-node
```

- This will show node assignment and pod status.

---

## 4. **List Istio CNI Pods and Their Nodes**

```bash
kubectl -n kube-system get pods -o wide | grep istio-cni
```

- Each node should have one running `istio-cni-node-xxxx` pod.

---

## 📚 References

- [Istio CNI Documentation](https://istio.io/latest/docs/setup/additional-setup/cni/)
- [Kubernetes DaemonSet Concepts](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/)

---
