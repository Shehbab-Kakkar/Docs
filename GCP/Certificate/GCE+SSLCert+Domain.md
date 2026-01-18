Below is a **clear, end-to-end guide** to using the **Google-managed Ingress controller (GCE Ingress)** on **GKE**.

This is the controller that automatically provisions:

* Google Cloud HTTP(S) Load Balancer
* Google-managed SSL certificates
* Backend services, health checks, forwarding rules, etc.

---

## 1️⃣ Prerequisites

Make sure you have:

* ✅ A **GKE cluster** (Standard or Autopilot)
* ✅ A **Service of type NodePort or ClusterIP**
* ✅ DNS A record pointing your domain → Ingress external IP
* ✅ HTTP(S) Load Balancing enabled (default in GKE)

Check controller availability:

```bash
kubectl get pods -n kube-system | grep ingress
```

---

## 2️⃣ Create a Service

Example backend service:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  type: NodePort
  selector:
    app: my-app
  ports:
  - port: 80
    targetPort: 8080
```

---

## 3️⃣ (Optional) Create a Google-managed certificate

```yaml
apiVersion: networking.gke.io/v1
kind: ManagedCertificate
metadata:
  name: my-managed-cert
spec:
  domains:
    - example.com
    - www.example.com
```

> ❗ No wildcards
> ❗ Max 100 domains per cert

Apply it:

```bash
kubectl apply -f managed-cert.yaml
```

---

## 4️⃣ Create an Ingress using the Google-managed controller

### Minimum required annotation

```yaml
kubernetes.io/ingress.class: "gce"
```

### Full example with SSL

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-ingress
  annotations:
    kubernetes.io/ingress.class: "gce"
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

Apply it:

```bash
kubectl apply -f ingress.yaml
```

---

## 5️⃣ Get the external IP

```bash
kubectl get ingress my-ingress
```

Example output:

```
ADDRESS         PORTS
34.120.xxx.xxx  80,443
```

➡️ Point your domain’s **A record** to this IP.

---

## 6️⃣ Verify certificate status

```bash
kubectl describe managedcertificate my-managed-cert
```

Look for:

```
Status: Active
```

Provisioning usually takes **10–60 minutes**.

---

## 7️⃣ Common annotations you may need

### Force HTTPS redirect

```yaml
ingress.kubernetes.io/ssl-redirect: "true"
```

### Use static IP

```yaml
kubernetes.io/ingress.global-static-ip-name: my-static-ip
```

### Multiple managed certs

```yaml
networking.gke.io/managed-certificates: cert1,cert2
```

---

## 8️⃣ What Google-managed Ingress creates for you

Automatically provisions:

* Global HTTP(S) Load Balancer
* Backend services
* Health checks
* URL maps
* SSL certs (if configured)

No manual GCP networking setup needed.

---

## 9️⃣ When NOT to use GCE Ingress

❌ If you need:

* Wildcard certs
* mTLS
* Advanced L7 routing
* NGINX-specific features

👉 Use **NGINX Ingress** or **Gateway API + Certificate Manager** instead.

---

If you want, I can:

* Help migrate from **NGINX → GCE Ingress**
* Show **Gateway API** (newer & recommended by Google)
* Debug **Ingress stuck in “PROVISIONING”**

Just tell me what you’re trying to achieve.
