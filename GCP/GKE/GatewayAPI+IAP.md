Below is a **concrete, end-to-end example of using Identity-Aware Proxy (IAP) with the Kubernetes Gateway API on GKE**.
This reflects the **current recommended pattern**: IAP is attached to a **Google Cloud HTTPS Load Balancer** that is **provisioned by the Gateway API**.

---

## Architecture overview

```
User
 └─> Google HTTPS Load Balancer
      └─ IAP (OAuth2 / Google Identity)
          └─ Gateway (Gateway API, GKE)
               └─ HTTPRoute
                    └─ Kubernetes Service
                         └─ Pods
```

Key points:

* **IAP is configured on the BackendService**, not inside Kubernetes
* **Gateway API** creates the load balancer and backend service
* Pods remain unaware of authentication
* IAP headers are injected for identity-aware apps

---

## Prerequisites

* GKE **Standard** or **Autopilot**
* Gateway API enabled:

  ```bash
  gcloud container clusters update CLUSTER_NAME \
    --gateway-api=standard \
    --location=REGION
  ```
* A **Google OAuth brand + client ID**
* DNS name pointing to the Gateway IP

---

## 1. Deploy a sample application

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello
spec:
  replicas: 1
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

```yaml
apiVersion: v1
kind: Service
metadata:
  name: hello
spec:
  selector:
    app: hello
  ports:
  - port: 80
    targetPort: 8080
```

---

## 2. Create a Gateway (managed HTTPS LB)

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: iap-gateway
spec:
  gatewayClassName: gke-l7-global-external-managed
  listeners:
  - name: https
    protocol: HTTPS
    port: 443
    tls:
      mode: Terminate
      certificateRefs:
      - kind: Secret
        name: tls-cert
```

> TLS cert can be from:
>
> * Google-managed cert (`ManagedCertificate`)
> * cert-manager
> * Kubernetes TLS secret

---

## 3. Create an HTTPRoute

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: hello-route
spec:
  parentRefs:
  - name: iap-gateway
  rules:
  - backendRefs:
    - name: hello
      port: 80
```

At this point:

* A **Global HTTPS Load Balancer** is created
* A **BackendService** exists in GCP

---

## 4. Enable IAP on the BackendService

### Find the backend service

```bash
gcloud compute backend-services list
```

Identify the one created by Gateway (name starts with `k8s2-`).

---

### Enable IAP

```bash
gcloud compute backend-services update BACKEND_SERVICE_NAME \
  --iap=enabled,oauth2-client-id=CLIENT_ID,oauth2-client-secret=CLIENT_SECRET \
  --global
```

---

## 5. Grant access to users

```bash
gcloud iap web add-iam-policy-binding \
  --resource-type=backend-services \
  --resource=BACKEND_SERVICE_NAME \
  --global \
  --member=user:alice@example.com \
  --role=roles/iap.httpsResourceAccessor
```

---

## 6. Access identity in your app (optional)

IAP injects headers such as:

```http
X-Goog-Authenticated-User-Email: accounts.google.com:alice@example.com
X-Goog-Authenticated-User-ID: 1234567890
```

Example (Node.js):

```js
const user = req.headers['x-goog-authenticated-user-email'];
```

⚠️ Always validate:

* Request comes from Google LB IPs
* Header exists (never trust blindly)

---

## Common pitfalls

| Issue            | Cause                                    |
| ---------------- | ---------------------------------------- |
| 403 from IAP     | User missing `iap.httpsResourceAccessor` |
| Redirect loop    | HTTP listener instead of HTTPS           |
| No headers       | IAP not enabled on backend service       |
| Cert not working | DNS not pointing to Gateway IP           |

---

## When to use this pattern

✅ Best for:

* Internal admin tools
* Zero-trust web apps
* Google Workspace users

❌ Not ideal for:

* Public APIs
* Non-browser OAuth flows

---

