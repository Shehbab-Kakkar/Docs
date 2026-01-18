In **GKE using the GCE (Google-managed) Ingress**, HTTP → HTTPS redirect is done with a **FrontendConfig** resource (this is the *correct and supported* way).

---

## ✅ Recommended way: `FrontendConfig` (GCE Ingress)

### 1️⃣ Create a FrontendConfig

```yaml
apiVersion: networking.gke.io/v1beta1
kind: FrontendConfig
metadata:
  name: https-redirect
spec:
  redirectToHttps:
    enabled: true
    responseCodeName: MOVED_PERMANENTLY_DEFAULT
```

Apply it:

```bash
kubectl apply -f frontendconfig.yaml
```

---

### 2️⃣ Attach it to your Ingress

Add this annotation to your **Ingress**:

```yaml
networking.gke.io/frontend-config: https-redirect
```

### Full Ingress example

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-ingress
  annotations:
    kubernetes.io/ingress.class: "gce"
    networking.gke.io/frontend-config: https-redirect
    networking.gke.io/managed-certificates: my-managed-cert
spec:
  rules:
  - host: example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: my-service
            port:
              number: 80
```

---

## 3️⃣ What this does

* 🔁 HTTP (port 80) → HTTPS (port 443)
* 🔒 Uses Google Cloud HTTPS Load Balancer
* ⚡ Redirect happens **at the load balancer** (fast, no pod traffic)
* 📈 SEO-safe (301 permanent redirect)

---

## 4️⃣ Verify it works

```bash
curl -I http://example.com
```

Expected response:

```
HTTP/1.1 301 Moved Permanently
Location: https://example.com/
```

---

## ❌ What NOT to use

| Method                               | Status                 |
| ------------------------------------ | ---------------------- |
| `ingress.kubernetes.io/ssl-redirect` | ❌ Not supported by GCE |
| App-level redirect                   | ❌ Inefficient          |
| NGINX annotations                    | ❌ Ignored by GCE       |

---

## 🧠 Notes / Gotchas

* Requires **GCE Ingress**, not NGINX
* Works with **Google-managed certificates**
* Takes **a few minutes** to propagate
* Supported in **Standard & Autopilot**

---

## 🆕 Alternative (newer): Gateway API

If you’re building new infrastructure, Google recommends:

* **Gateway API + Certificate Manager**
* Native HTTPS redirects built into `HTTPRoute`

I can show that setup too if you want.

---

If your redirect **isn’t working**, tell me:

* GKE version
* Autopilot or Standard
* Output of `kubectl describe ingress`

I’ll help you debug it quickly 👍
