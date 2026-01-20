Below is a **minimal, working example** of **GKE Multi-Cluster Ingress (MCI)** YAML using the **Kubernetes-native approach**.

This assumes:

* You have **2+ regional GKE clusters**
* All clusters are **registered to the same Fleet**
* Same namespaces and services exist in each cluster
* You’re using the **GKE Multi-Cluster Ingress controller** (not deprecated Anthos Ingress)

---

## 1️⃣ Backend Service (deployed to *every* cluster)

```yaml
apiVersion: v1
kind: Service
metadata:
  name: hello-service
  namespace: default
spec:
  selector:
    app: hello
  ports:
    - port: 80
      targetPort: 8080
  type: ClusterIP
```

---

## 2️⃣ Deployment (deployed to *every* cluster)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello
  namespace: default
spec:
  replicas: 3
  selector:
    matchLabels:
      app: hello
  template:
    metadata:
      labels:
        app: hello
    spec:
      containers:
        - name: hello
          image: gcr.io/google-samples/hello-app:1.0
          ports:
            - containerPort: 8080
```

---

## 3️⃣ Multi-Cluster Ingress (apply **once**, to *any* cluster)

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: hello-mci
  namespace: default
  annotations:
    kubernetes.io/ingress.class: "gce-multi-cluster"
spec:
  rules:
    - host: hello.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: hello-service
                port:
                  number: 80
```

✅ This single Ingress creates:

* A **global HTTP(S) Load Balancer**
* Backends in **all clusters**
* Geo-based routing to nearest healthy region

---

## 4️⃣ Optional: ManagedCertificate (TLS)

```yaml
apiVersion: networking.gke.io/v1
kind: ManagedCertificate
metadata:
  name: hello-cert
  namespace: default
spec:
  domains:
    - hello.example.com
```

Attach it to the Ingress:

```yaml
metadata:
  annotations:
    networking.gke.io/managed-certificates: hello-cert
```

---

## 5️⃣ (Optional) BackendConfig for health checks & timeouts

```yaml
apiVersion: cloud.google.com/v1
kind: BackendConfig
metadata:
  name: hello-backendconfig
  namespace: default
spec:
  timeoutSec: 30
  healthCheck:
    requestPath: /
    port: 8080
```

Attach to the Service:

```yaml
metadata:
  annotations:
    cloud.google.com/backend-config: '{"default": "hello-backendconfig"}'
```

---

## Important notes (common gotchas)

### ✅ Required

* Same **service name + namespace** in all clusters
* Same **port numbers**
* Clusters must be in the **same Fleet**
* Ingress applied **once only**

### ❌ Not allowed

* NodePort services
* Different selectors per cluster
* Different namespaces per cluster

---

## How traffic flows

1. User hits global Anycast IP
2. Google Front Ends route to nearest healthy region
3. Health checks determine active backends
4. Automatic regional failover

---

## When to use this

✔ Stateless services
✔ Active-active multi-region
✔ Simple global routing

## When *not* to use it

✘ Complex L7 routing rules
✘ Custom Envoy/Istio needs
✘ Per-region custom domains

---



* Convert this to **Terraform**
* Add **canary traffic splitting**
* Show **MultiClusterService** (Service Export/Import) example
* Compare with **Anthos Service Mesh**

Just tell me 👍
