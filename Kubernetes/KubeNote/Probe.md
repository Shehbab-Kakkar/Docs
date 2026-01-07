Here’s a **clear comparison** of Readiness Probe, Liveness Probe, and Startup Probe in Kubernetes:

---

### **1. Readiness Probe**
- **Purpose:**  
  Checks if the container is **ready to accept traffic**.
- **Action if Fails:**  
  The pod is marked as "not ready"; it will be **removed from Service load balancers** and **not receive requests** until it passes.
- **Typical Use:**  
  When your app needs some initialization time before it can serve requests (e.g., wait for DB connection).

---

### **2. Liveness Probe**
- **Purpose:**  
  Checks if the container is **still alive/healthy** (not stuck or deadlocked).
- **Action if Fails:**  
  **Container is restarted** by kubelet (self-healing).
- **Typical Use:**  
  To recover from deadlocks, application crashes, or unresponsiveness.

---

### **3. Startup Probe**
- **Purpose:**  
  Designed for **slow-starting containers**—checks if your app has **successfully started**.
- **Action if Fails:**  
  Same as liveness probe: container will be restarted.
- **Special Behavior:**  
  **While startup probe is running, liveness and readiness probes are *ignored.***  
  Once startup probe succeeds, liveness and readiness probes take over.
- **Typical Use:**  
  Useful for apps that need thorough initialization before they can be considered “alive” or “ready.”

---

### **Summary Table**

| Probe            | Checks If...                 | Failure Action                          | When to Use                      |
|------------------|-----------------------------|-----------------------------------------|----------------------------------|
| Readiness Probe  | Ready to accept traffic      | Removed from Service endpoints          | For initializing apps            |
| Liveness Probe   | Still healthy (not stuck)    | Container restarted                     | For deadlocks/crashes            |
| Startup Probe    | Has successfully started     | Container restarted (probes ignored until success) | For slow-starting apps           |

---

**Example Usage in Pod spec:**
```yaml
readinessProbe:
  httpGet:
    path: /healthz
    port: 8080
  initialDelaySeconds: 5

livenessProbe:
  httpGet:
    path: /live
    port: 8080
  initialDelaySeconds: 15

startupProbe:
  httpGet:
    path: /startup
    port: 8080
  failureThreshold: 30
  periodSeconds: 10
```

---

**In Short:**  
- **Readiness Probe:** Controls traffic flow to pod.  
- **Liveness Probe:** Triggers restarts if unhealthy.  
- **Startup Probe:** Gives slow apps extra time before other probes start.
