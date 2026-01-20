Below is a **complete, end-to-end GKE configuration** that includes:

✅ GKE cluster
✅ Deployment (pods)
✅ Service
✅ Ingress (GCE)
✅ **Cloud CDN enabled**
✅ **Managed SSL certificate**
✅ CLI commands + YAML

This is the **standard, production-ready GKE + Cloud CDN setup**.

---

# 1️⃣ Create GKE cluster (CLI)

```bash
gcloud container clusters create web-cluster \
  --region us-central1 \
  --num-nodes 3 \
  --enable-ip-alias \
  --enable-autoupgrade \
  --enable-autorepair
```

Get credentials:

```bash
gcloud container clusters get-credentials web-cluster \
  --region us-central1
```

---

# 2️⃣ Deployment (Pods)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: web
  template:
    metadata:
      labels:
        app: web
    spec:
      containers:
        - name: web
          image: gcr.io/google-samples/hello-app:1.0
          ports:
            - containerPort: 8080
```

Apply:

```bash
kubectl apply -f deployment.yaml
```

---

# 3️⃣ Service (ClusterIP)

```yaml
apiVersion: v1
kind: Service
metadata:
  name: web-service
  annotations:
    cloud.google.com/backend-config: '{"default": "cdn-backendconfig"}'
spec:
  type: ClusterIP
  selector:
    app: web
  ports:
    - port: 80
      targetPort: 8080
```

Apply:

```bash
kubectl apply -f service.yaml
```

---

# 4️⃣ BackendConfig (Enable Cloud CDN)

This is **where Cloud CDN is enabled**.

```yaml
apiVersion: cloud.google.com/v1
kind: BackendConfig
metadata:
  name: cdn-backendconfig
spec:
  cdn:
    enabled: true
    cachePolicy:
      includeHost: true
      includeProtocol: true
      includeQueryString: true
```

Apply:

```bash
kubectl apply -f backendconfig.yaml
```

---

# 5️⃣ Managed SSL Certificate

```yaml
apiVersion: networking.gke.io/v1
kind: ManagedCertificate
metadata:
  name: web-cert
spec:
  domains:
    - example.com
```

⚠️ Make sure:

* `example.com` DNS A record points to the **Ingress IP**

Apply:

```bash
kubectl apply -f certificate.yaml
```

---

# 6️⃣ Ingress (GCE + HTTPS + Cloud CDN)

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: web-ingress
  annotations:
    kubernetes.io/ingress.class: "gce"
    networking.gke.io/managed-certificates: web-cert
    kubernetes.io/ingress.allow-http: "false"
spec:
  rules:
    - host: example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: web-service
                port:
                  number: 80
```

Apply:

```bash
kubectl apply -f ingress.yaml
```

---

# 7️⃣ Get Ingress IP

```bash
kubectl get ingress web-ingress
```

Output example:

```text
ADDRESS         34.120.xxx.xxx
```

➡️ Point your DNS **A record** to this IP.

---

# 8️⃣ Verify Cloud CDN is enabled

```bash
gcloud compute backend-services list
```

Find your backend, then:

```bash
gcloud compute backend-services describe BACKEND_NAME \
  --global
```

Look for:

```yaml
cdnPolicy:
  enabled: true
```

---

# 9️⃣ Request flow (important)

```
User
 ↓
Cloud CDN (edge cache)
 ↓
Global HTTP(S) Load Balancer
 ↓
GKE Ingress
 ↓
Service (ClusterIP)
 ↓
Pods
```

---

# 10️⃣ Common production hardening (recommended)

✔ Add readiness/liveness probes
✔ Set resource requests/limits
✔ Enable HPA
✔ Use private GKE cluster
✔ Enable logging & monitoring

---

# 🧠 Key Notes (very important)

| Component  | Requirement                     |
| ---------- | ------------------------------- |
| Cloud CDN  | Only works with **GCE Ingress** |
| Service    | Must be **ClusterIP**           |
| TLS        | Use **ManagedCertificate**      |
| Caching    | App must send cacheable headers |
| Ingress IP | DNS must point to it            |

---

## Want next?

I can:

* Add **HPA + autoscaling**
* Convert everything to **Terraform**
* Add **multi-region + Cloud CDN**
* Add **IAP / Cloud Armor / WAF**

Just tell me 👍
