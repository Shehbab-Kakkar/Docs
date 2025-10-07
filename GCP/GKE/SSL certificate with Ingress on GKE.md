Here's a **complete example** of setting up an **SSL certificate with Ingress on GKE (Google Kubernetes Engine)** using **HTTPS** via **Google-managed SSL certificates**.

---

## ✅ Requirements

* GKE cluster running
* Domain name (e.g., `example.com`) pointing to your Ingress IP
* `kubectl` and `gcloud` CLI installed and authenticated

---

## 🔧 Step-by-Step: GKE Ingress with SSL Certificate

### 1. **Create a Managed SSL Certificate**

```yaml
# ssl-cert.yaml
apiVersion: networking.gke.io/v1
kind: ManagedCertificate
metadata:
  name: my-ssl-cert
spec:
  domains:
    - example.com
```

Apply it:

```bash
kubectl apply -f ssl-cert.yaml
```

---

### 2. **Create a Deployment and Service**

```yaml
# app.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-app
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
      - name: hello-app
        image: gcr.io/google-samples/hello-app:1.0
        ports:
        - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: hello-service
spec:
  selector:
    app: hello
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  type: NodePort
```

Apply it:

```bash
kubectl apply -f app.yaml
```

---

### 3. **Create the Ingress Resource with TLS**

```yaml
# ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: hello-ingress
  annotations:
    kubernetes.io/ingress.class: "gce"  # For GKE Ingress
    networking.gke.io/managed-certificates: my-ssl-cert
spec:
  rules:
  - host: example.com
    http:
      paths:
      - pathType: Prefix
        path: "/"
        backend:
          service:
            name: hello-service
            port:
              number: 80
```

Apply it:

```bash
kubectl apply -f ingress.yaml
```

---

### 4. **Wait for the Ingress to Get an IP**

```bash
kubectl get ingress hello-ingress
```

Point your domain (`example.com`) to the **EXTERNAL-IP** shown above (create an A record in your DNS).

---

### 5. **Check the Certificate Status**

```bash
kubectl describe managedcertificate my-ssl-cert
```

Look for:

```
Status:
  CertificateStatus: Active
```

It may take a few minutes for Google to provision the SSL certificate.

---

## ✅ Done!

Now your app should be available at:

```
https://example.com
```

With a valid **Google-managed SSL certificate**.

---

