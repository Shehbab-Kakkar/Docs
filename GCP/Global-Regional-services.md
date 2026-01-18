Below is a **clear, practical list of Google Cloud services by scope**.
(I’ve grouped the most commonly used ones — not every niche product, but the ones architects and engineers rely on.)

---

## 🌍 **Global services (not tied to a region)**

These services work **across all regions** or are managed globally.

### **Networking**

* **VPC Network**
* **Firewall rules**
* **Routes**
* **Cloud Load Balancing (Global HTTP(S))**
* **Cloud CDN**
* **Cloud DNS**
* **Network Service Tiers**
* **Shared VPC**

### **Compute & Platform**

* **App Engine (standard & flexible – global routing)**
* **Cloud Run (service is regional, but domain + routing can be global)**
* **Anthos**
* **Cloud Functions (2nd gen – managed globally, runs regionally)**

### **Security & Identity**

* **IAM**
* **Cloud Identity**
* **Organization / Folder / Project**
* **Secret Manager**
* **Cloud KMS (keys are regional, policies are global)**

### **Operations & Management**

* **Cloud Monitoring**
* **Cloud Logging**
* **Cloud Trace**
* **Cloud Profiler**
* **Cloud Error Reporting**
* **Billing**
* **Resource Manager**

---

## 📍 **Regional services**

These resources are created **in a specific region**.

### **Networking**

* **Subnets**
* **Regional Load Balancers**

  * Internal TCP/UDP Load Balancer
  * Internal HTTP(S) Load Balancer
* **Cloud NAT**
* **Cloud Router**
* **VPN Gateway (HA VPN)**
* **Interconnect attachments**

### **Compute**

* **Compute Engine regional managed instance groups**
* **Cloud Run services**
* **Cloud Functions**
* **App Engine services (runtime runs in a region)**

### **Data & Storage**

* **Cloud SQL**
* **Cloud Spanner (regional or multi-regional configs)**
* **Bigtable**
* **Memorystore (Redis / Memcached)**
* **Filestore**
* **Persistent Disks (regional PD)**

### **Analytics & Messaging**

* **Pub/Sub (topics are global, but data placement is regional)**
* **Dataflow**
* **Dataproc**
* **Composer (Airflow)**

---

## 📦 **Zonal services (for completeness)**

These live **inside a zone** (which belongs to a region).

* **Compute Engine VM instances**
* **Zonal Persistent Disks**
* **Zonal Managed Instance Groups**
* **GPUs / TPUs**
* **Local SSDs**

---

## 🧠 Quick memory aid

* **VPC = Global**
* **Subnet = Regional**
* **VM = Zonal**
* **Firewall = Global**
* **Load balancer = Global or Regional (depends on type)**

---

If you want, I can:

* Give you a **one-page architecture cheat sheet**
* Explain **why some services are global vs regional**
* Map these to **AWS equivalents**
* Help design a **multi-region architecture**

Just tell me what you’re building.
