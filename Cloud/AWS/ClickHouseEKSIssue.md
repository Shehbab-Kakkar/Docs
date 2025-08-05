# Debugging Distributed Databases: Beyond the Stack

## 🐞 The Problem

While running **ClickHouse** on a Kubernetes cluster in AWS, everything kept failing. But it wasn’t the database.

- **No query spikes.**
- **No errors in logs.**
- **Just pods endlessly crashing, sometimes taking down whole nodes.**

### Attempts to Fix

- **Scaled the node group**: Added more nodes.
- **Result:** The more nodes I added, the faster they filled. The pod crashes continued.
- **Checked compute and memory:** No shortages.
- **Checked storage:** Healthy.
- **Checked distribution:** Pods and services were balanced.

> **It didn’t look like a resource problem. But it was. Just not the one I expected.**

---

## 🔍 Root Cause Analysis

### **1. Pod IP Exhaustion**

- **Kubernetes on AWS*** assigns pod IPs based on each node’s Elastic Network Interface (ENI) capacity.
- On `t3.medium`, only **17 pods** could get IPs per node by default.
- When the IP limit was reached, further pods wouldn't schedule, leading to instability.

#### **Resolution #1: Expanded Pod IP Allocation—Prefix Delegation**
- **Action:** Enabled _prefix delegation_ in the AWS VPC CNI.
- **Effect:** Increased the per-node IP allocation from **17 to 117**, unblocking pod scheduling.
- **But —** crashes kept coming.

---

### **2. ENI Throughput Saturation**

- Used [**Pixie**](https://pixielabs.ai/) (an observability tool): revealed no stable pods, no traces, and lots of static.
- **Root Cause:** The t3.medium’s ENI reached max bandwidth and packets-per-second (PPS).
- **Pods failed** not because of database errors, but due to saturated network limits at the instance and ENI level.

#### **Resolution #2: Upgrade Instance Type**
- **Action:** Switched to larger EC2 instance types **(with higher network and ENI throughput)**.
- **Result:** Pods stabilized, ClickHouse scaled, Pixie monitoring revived.

---

## ⚡️ Key Takeaways

- **Scaling isn’t just horizontal** (adding more nodes). Sometimes, it’s vertical (beefier nodes) and _deeper_ (network, kernel, CNI settings).
- **Pod scheduling limits** can be non-obvious: not all resource issues announce themselves via CPU/RAM metrics.
- **Cloud-native environments introduce new limits** (e.g., ENI throughput, pod IP addresses).
- **Tools like Pixie reveal what logs can’t:** bottlenecks beneath your application stack.

> **We spend years tuning queries and configs, but real resilience starts when you debug beneath the stack.  
> Scaling doesn’t just mean wider — it’s also about the cloud plumbing below.**

---

### 📊 Example: Healthy Cluster (Pixie Flamegraph Visualization)

![Pixie Flamegraph Example](example-pixie-flamegraph.png)

---

## References

- [AWS VPC CNI Prefix Delegation Documentation](https://docs.aws.amazon.com/eks/latest/userguide/pod-networking.html)
- [AWS EC2 Instance Network Performance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#instance-type-networking)

*For issues, contributions, or more stories from production: open an Issue or PR!*
