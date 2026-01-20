Below is a **complete add-on configuration** for your existing GKE + Ingress + Cloud CDN setup that adds:

✅ **IAP (Identity-Aware Proxy)**
✅ **Cloud Armor (WAF)**
✅ **HPA (Horizontal Pod Autoscaler)**
✅ **Cluster autoscaling**

I’ll show **what must be YAML vs CLI** (some things cannot be done purely in YAML).

---

# 1️⃣ Enable required APIs (CLI)

```bash
gcloud services enable \
  iap.googleapis.com \
  compute.googleapis.com \
  container.googleapis.com \
  cloudresourcemanager.googleapis.com
```

---

# 2️⃣ HPA – Horizontal Pod Autoscaler (YAML)

## Deployment (add resources first)

Update your existing `Deployment`:

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
          resources:
            requests:
              cpu: "100m"
              memory: "128Mi"
            limits:
              cpu: "500m"
              memory: "256Mi"
```

Apply:

```bash
kubectl apply -f deployment.yaml
```

---

## HPA YAML

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: web-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: web-app
  minReplicas: 3
  maxReplicas: 10
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 70
```

Apply:

```bash
kubectl apply -f hpa.yaml
```

Verify:

```bash
kubectl get hpa
```

---

# 3️⃣ Cluster Autoscaling (CLI)

If your cluster **already exists**, enable autoscaling on the node pool:

```bash
gcloud container clusters update web-cluster \
  --region us-central1 \
  --enable-autoscaling \
  --min-nodes 3 \
  --max-nodes 10
```

Check:

```bash
kubectl get nodes
```

---

# 4️⃣ Cloud Armor (WAF)

## Create a Cloud Armor security policy (CLI)

```bash
gcloud compute security-policies create web-waf \
  --description "WAF for GKE Ingress"
```

### Add OWASP-style rules

```bash
gcloud compute security-policies rules create 1000 \
  --security-policy web-waf \
  --expression "evaluatePreconfiguredWaf('xss-v33-stable')" \
  --action deny-403

gcloud compute security-policies rules create 1001 \
  --security-policy web-waf \
  --expression "evaluatePreconfiguredWaf('sqli-v33-stable')" \
  --action deny-403
```

---

## Attach Cloud Armor to GKE via BackendConfig (YAML)

Update your existing `BackendConfig`:

```yaml
apiVersion: cloud.google.com/v1
kind: BackendConfig
metadata:
  name: cdn-backendconfig
spec:
  cdn:
    enabled: true
  securityPolicy:
    name: web-waf
```

Apply:

```bash
kubectl apply -f backendconfig.yaml
```

---

# 5️⃣ IAP (Identity-Aware Proxy)

⚠️ **Important limitation**

> IAP **cannot be fully configured via YAML**.
> It must be enabled on the **HTTP(S) Load Balancer backend service**.

---

## Step 1: Create OAuth consent screen (one-time)

```bash
gcloud iap oauth-brands create \
  --application_title="GKE App" \
  --support_email=YOUR_EMAIL
```

---

## Step 2: Create OAuth client

```bash
gcloud iap oauth-clients create \
  --display_name="gke-iap-client"
```

Save:

* `CLIENT_ID`
* `CLIENT_SECRET`

---

## Step 3: Enable IAP on backend service

Find backend service:

```bash
gcloud compute backend-services list --global
```

Enable IAP:

```bash
gcloud compute backend-services update BACKEND_NAME \
  --global \
  --iap=enabled,oauth2-client-id=CLIENT_ID,oauth2-client-secret=CLIENT_SECRET
```

---

## Step 4: Grant access to users

```bash
gcloud iap web add-iam-policy-binding \
  --resource-type=backend-services \
  --resource-id=BACKEND_NAME \
  --member=user:user@example.com \
  --role=roles/iap.httpsResourceAccessor
```

---

# 6️⃣ Final Request Flow

```
User
 ↓
Cloud Armor (WAF)
 ↓
Cloud CDN (cached if possible)
 ↓
IAP (auth)
 ↓
Global HTTPS Load Balancer
 ↓
GKE Ingress
 ↓
Service
 ↓
Pods (HPA)
```

---

# 7️⃣ Verification checklist

### HPA

```bash
kubectl get hpa
kubectl top pods
```

### Cloud Armor

```bash
gcloud compute security-policies list
```

### IAP

* Open site in browser → Google login prompt
* Access denied for unauthorized users

---

# 8️⃣ Production best practices

✔ Use **private GKE cluster**
✔ Enable **Cloud Logging + Monitoring**
✔ Add **Cloud Armor rate limiting**
✔ Use **separate backend services per app**
✔ Enable **HTTPS only**

---

## Want next?

I can:

* Add **rate limiting rules**
* Add **JWT-based IAP auth for APIs**
* Convert everything to **Terraform**
* Add **multi-region failover**

Just tell me what’s next 🚀
