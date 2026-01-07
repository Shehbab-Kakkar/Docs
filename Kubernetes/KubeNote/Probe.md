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


Certainly! Here’s a breakdown of the common probe parameters in Kubernetes (applies to **readinessProbe, livenessProbe, and startupProbe**):

---

## 1. **Common Probe Parameters**

| Parameter             | Description                                                                                  |
|-----------------------|----------------------------------------------------------------------------------------------|
| **initialDelaySeconds** | Seconds to wait after container starts before first probe is initiated (default: 0)         |
| **periodSeconds**       | How often (in seconds) to perform the probe (default: 10)                                   |
| **timeoutSeconds**      | Number of seconds after which the probe times out (default: 1)                              |
| **successThreshold**    | Minimum consecutive successes for probe to be considered successful (default: 1)            |
| **failureThreshold**    | When a probe fails, number of retries before container is restarted or pod is marked unready (default: 3) |

---

## 2. **Probe Types**

You can only choose _one_ type per probe. The main types are:

### HTTP GET probe
```yaml
httpGet:
  path: /healthz
  port: 8080
  httpHeaders:
  - name: Custom-Header
    value: Awesome
  scheme: HTTP
```
- **path**: Endpoint path to check (e.g., `/healthz`)
- **port**: Container port to call
- **httpHeaders**: Custom HTTP headers
- **scheme**: `HTTP` or `HTTPS`

---

### TCP Socket probe
```yaml
tcpSocket:
  port: 8080
```
- **port**: Port to check for an open TCP connection

---

### Exec probe
```yaml
exec:
  command:
    - cat
    - /tmp/healthy
```
- **command**: Command to run inside the container; success if exit code = 0

---

## 3. **All Parameters—Explanation Table**

| Parameter              | Applies to       | Description                                                           |
|------------------------|------------------|-----------------------------------------------------------------------|
| initialDelaySeconds    | All probes       | Wait time in seconds before performing first check after container start |
| periodSeconds          | All probes       | Time between each probe attempt (seconds)                             |
| timeoutSeconds         | All probes       | How long to wait for a response before it’s considered a failure      |
| successThreshold       | All probes       | How many successes needed to mark container healthy/ready             |
| failureThreshold       | All probes       | Failures allowed before considered failed/unready (or restarted)      |
| httpGet                | HTTP/TCP probes  | Parameters for HTTP GET requests (path, port, headers, scheme)        |
| tcpSocket              | TCP probes       | Parameters for checking TCP (port)                                    |
| exec                   | Exec probes      | Command to execute for health check                                   |

---

## **Example Probe Spec**

```yaml
livenessProbe:
  httpGet:
    path: /healthz
    port: 8080
  initialDelaySeconds: 10
  periodSeconds: 5
  timeoutSeconds: 2
  successThreshold: 1
  failureThreshold: 3
```

---

### **In summary:**
- Use these parameters to **tune how and when Kubernetes checks container health** and what happens when it fails.
- The flexibility lets you accommodate fast, slow, or unreliable startup and health patterns!
