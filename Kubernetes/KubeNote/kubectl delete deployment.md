Short answer: **`kubectl delete deployment` doesn’t directly kill pods** — it tells Kubernetes to remove the *Deployment object*, and Kubernetes cleans up everything under it in a controlled way.

Here’s what actually happens under the hood 👇

---

## 1. `kubectl` talks to the API server

When you run:

```bash
kubectl delete deployment my-deploy
```

`kubectl`:

* Authenticates using your kubeconfig
* Sends an HTTP **DELETE** request to the Kubernetes **API Server**
* The request targets the Deployment resource in etcd

`kubectl` itself doesn’t delete anything — it’s just a client.

---

## 2. Deployment is marked for deletion

The API server:

* Sets a **`deletionTimestamp`** on the Deployment
* The Deployment enters a **“terminating”** state
* Actual removal is delayed until cleanup finishes

This is important: deletion is **asynchronous**.

---

## 3. Cascading deletion kicks in (default behavior)

By default, Kubernetes uses **cascading deletion** (`--cascade=background`).

So Kubernetes:

* Finds resources **owned by the Deployment**
* Uses **OwnerReferences** to track relationships

Ownership chain looks like this:

```
Deployment
 └── ReplicaSet
      └── Pods
```

---

## 4. ReplicaSets are deleted

The Deployment controller:

* Deletes the ReplicaSet(s) it created
* Each ReplicaSet deletion triggers pod cleanup

If you run:

```bash
kubectl get rs
```

You’ll see them disappearing before the pods are fully gone.

---

## 5. Pods are gracefully terminated

For each Pod:

* Pod gets `deletionTimestamp`
* Kubelet sends **SIGTERM** to containers
* Kubernetes waits for `terminationGracePeriodSeconds` (default: 30s)
* If the container doesn’t exit → **SIGKILL**

This is why pods often stay in `Terminating` for a bit.

---

## 6. Finalizers (if any) must complete

If the Deployment, ReplicaSet, or Pods have **finalizers**:

* Kubernetes waits for them to finish
* Only then is the object fully removed from etcd

Finalizers are common with:

* Cloud load balancers
* Storage (PVC cleanup)
* Custom controllers

---

## 7. Objects are removed from etcd

Once:

* All child resources are gone
* Finalizers are cleared

Kubernetes:

* Deletes the Deployment record from etcd
* The resource fully disappears

---

## Useful flags you might see

```bash
kubectl delete deployment my-deploy --cascade=foreground
```

* Waits until **all child resources are deleted first**

```bash
kubectl delete deployment my-deploy --cascade=orphan
```

* Deletes Deployment
* **Leaves ReplicaSets and Pods running**

```bash
kubectl delete deployment my-deploy --grace-period=0 --force
```

* Immediate deletion (dangerous; skips graceful shutdown)

---

## TL;DR

* `kubectl` sends a DELETE request
* API server marks Deployment for deletion
* Kubernetes follows owner references
* ReplicaSets → Pods are deleted
* Pods shut down gracefully
* Everything is finally removed from etcd

If you want, I can also explain this **from the controller’s perspective**, or show the **exact API calls** involved.
