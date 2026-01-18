For **Google-managed certificates with GKE Ingress**, you enable them via a **Kubernetes annotation** on the Ingress and a **ManagedCertificate** resource.

Below is the standard and correct setup.

---

## 1️⃣ ManagedCertificate resource

Create a `ManagedCertificate` object listing your domains (non-wildcard only):

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

> ⚠️ Max **100 non-wildcard domains** per ManagedCertificate.

---

## 2️⃣ Ingress annotation (this is what you asked for)

Add the following annotation to your **Ingress**:

```yaml
networking.gke.io/managed-certificates: my-managed-cert
```

### Full Ingress example

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

---

## 3️⃣ Multiple certificates (optional)

If you need more than 100 domains, you can attach **multiple managed certs**:

```yaml
annotations:
  networking.gke.io/managed-certificates: cert-1,cert-2
```

---

## 4️⃣ Important notes / gotchas

* ✅ Only works with **GCE Ingress** (not NGINX)
* ❌ Wildcards (`*.example.com`) are **not supported**
* 🔄 Certificate provisioning can take **10–60 minutes**
* 🔍 Check status with:

  ```bash
  kubectl describe managedcertificate my-managed-cert
  ```

---

If you want, I can also:

* Compare this with **Certificate Manager + Ingress**
* Show how to migrate from **pre-shared certs**
* Help debug a cert stuck in `PROVISIONING`

Just tell me 👍
