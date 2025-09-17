When a **MongoDB StatefulSet pod** is running in a **Kubernetes cluster**, and the **worker node hosting that pod goes down**, here's what happens:

---

### 🧠 1. **Kubernetes Will Notice the Node is Not Ready**

* The **kubelet** on the control plane monitors node health.
* When the node goes **NotReady** (e.g., due to power off or crash), Kubernetes marks it as **unavailable**.
* This can take **a few minutes** depending on `node-monitor-grace-period` (default: 40 seconds) and `pod-eviction-timeout` (default: 5 minutes).

---

### 🧱 2. **StatefulSet Pod Won't Automatically Move**

* **StatefulSet pods are bound to persistent identities** (like `mongodb-0`, `mongodb-1`, etc.), and do **not** get rescheduled to a new node automatically **while the old node is still marked as existing**, even if it's NotReady.
* This is **by design**, to maintain **identity and data consistency**, especially important for databases like MongoDB.

---

### ⚙️ 3. **What Will Happen to the Pod?**

* The pod (e.g., `mongodb-1`) will stay in a **"Terminating"** or **"Unknown"** state.
* It will **not be rescheduled** to another node **until**:

  * The failed node is manually removed or deleted from the cluster.
  * Or the pod is force-deleted and recreated manually on another node (with access to the same PVC).

---

### 🔧 4. **Recovery Options**

You have a few options to recover the pod:

#### ✅ **Option A: Node Comes Back**

* If the worker node is restarted and rejoins the cluster:

  * The MongoDB pod should resume operation normally (assuming volume is intact).

#### ❌ **Option B: Node is Permanently Dead**

If the node is permanently lost, you can:

1. **Delete the pod manually**:

   ```bash
   kubectl delete pod mongodb-1 --force --grace-period=0
   ```
2. **Ensure the PersistentVolumeClaim (PVC)** used by the pod is **on shared or networked storage** (like EBS, Ceph, etc.) and is available to other nodes.
3. Kubernetes will **reschedule** the pod to another available node in the cluster, attaching the existing PVC (if storage allows).

---

### 🔄 5. **High Availability Impact**

If you're using **MongoDB Replica Set** with multiple replicas (`mongodb-0`, `mongodb-1`, `mongodb-2`):

* The replica set can continue functioning **as long as a majority of members are alive**.
* If the downed pod was a **secondary**, MongoDB can tolerate the loss.
* If the downed pod was the **primary**, the replica set will **elect a new primary** if quorum is maintained.

---

### ✅ Summary

| Event               | Outcome                                           |
| ------------------- | ------------------------------------------------- |
| Node goes down      | Pod stays in "Unknown"/"Terminating" state        |
| Pod not rescheduled | Unless node is deleted or pod is force-deleted    |
| Data safe?          | Yes, if using persistent volume on shared storage |
| MongoDB impact?     | Minimal if quorum is preserved                    |

---

If the **worker node comes back online**, here's what happens to the **MongoDB pod** (part of a StatefulSet) and how it rejoins the **MongoDB replica set** cluster:

---

## ✅ Step-by-Step: What Happens When the Node Comes Back

### 1. **Node Status Changes to Ready**

* Kubernetes detects that the node is back.
* The node transitions from `NotReady` to `Ready`.

```bash
kubectl get nodes
```

You'll see the node is now `Ready`.

---

### 2. **Pod on That Node Is Revived**

* Since the pod was **never deleted**, it is still **bound to that node**.
* The kubelet on the recovered node restarts the pod from where it left off.
* If it had a **PersistentVolumeClaim (PVC)**, it will reattach and mount the volume.

```bash
kubectl get pod mongodb-1 -o wide
```

You should see it's running on the original node again.

---

### 3. **MongoDB Process Starts**

* The MongoDB process (`mongod`) inside the container starts.
* It reads the replica set config stored in the data directory.
* It **automatically reconnects to the replica set** as a member using the stored configuration.

MongoDB uses the internal replica set protocol to:

* Sync any missed oplog entries.
* Transition into `SECONDARY` or `PRIMARY` depending on the election state.

You can check the replica set status:

```bash
kubectl exec -it mongodb-0 -- mongo
> rs.status()
```

You should see the recovered pod (e.g., `mongodb-1`) listed with a healthy state like:

```json
{
  "name": "mongodb-1.mongodb.default.svc.cluster.local:27017",
  "stateStr": "SECONDARY",
  ...
}
```

---

## 🔁 What MongoDB Does Internally

Once the pod restarts:

* If it was the **primary** before the failure and another node took over, it will become a **secondary**.
* It will **sync the oplog** to catch up with the current state.
* No manual intervention is usually needed unless:

  * Data files are corrupted.
  * There’s a network partition.
  * Replica set config is out of sync.

---

## ⚠️ What You Need to Ensure

| Requirement                  | Why It Matters                                                                   |
| ---------------------------- | -------------------------------------------------------------------------------- |
| **Persistent Storage (PVC)** | So MongoDB retains its data after restart.                                       |
| **Stable Network Identity**  | StatefulSet ensures same hostname/IP is used (important for replica set config). |
| **Correct Readiness Probes** | Prevent MongoDB from being added to service endpoints too early.                 |

---

## 🛠️ Optional: Manual Checks (Not Usually Needed)

If something looks off (e.g., pod is not in replica set), you can manually re-add it:

```js
rs.add("mongodb-1.mongodb.default.svc.cluster.local:27017")
```

But **in a properly configured StatefulSet**, this is not required — the pod re-joins automatically.

---

## ✅ In Short:

| Action                    | Result                                   |
| ------------------------- | ---------------------------------------- |
| Node comes back           | Kubernetes sees it as Ready              |
| Pod restarts on that node | Uses same PVC and hostname               |
| MongoDB process restarts  | Reconnects to replica set                |
| Replica Set               | Syncs oplog and resumes normal operation |

---

