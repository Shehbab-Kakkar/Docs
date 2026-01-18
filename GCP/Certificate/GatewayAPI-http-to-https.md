 GKE HTTP-to-HTTPS redirection examples, including **both Ingress with FrontendConfig** and **Gateway API**. 
---

## **1. Using Ingress with FrontendConfig**

### **Step 1: Create a FrontendConfig**

```yaml
apiVersion: networking.gke.io/v1beta1
kind: FrontendConfig
metadata:
  name: http-to-https-redirect
spec:
  redirectToHttps:
    enabled: true
    responseCodeName: MOVED_PERMANENTLY_DEFAULT # Optional: 301 (default) or 308 (PERMANENT_REDIRECT)
```

---

### **Step 2: Update Your Ingress**

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-ingress
  annotations:
    kubernetes.io/ingress.class: "gce"
    networking.gke.io/v1beta1.FrontendConfig: "http-to-https-redirect"
spec:
  # Ensure HTTP is allowed so redirect can work (default is true)
  # kubernetes.io/ingress.allow-http: "true"
  tls:
    - secretName: my-tls-secret
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

**Key Points:**

* **Allow HTTP:** Do not set `kubernetes.io/ingress.allow-http: "false"`; Ingress must receive traffic on port 80 to issue the redirect.
* **Static IP:** Recommended to reserve a static IP for stable redirection.
* **Managed Certificates:** Works the same with Google-managed SSL certificates.
* **Alternative:** Console users can enable HTTP-to-HTTPS redirect via the frontend configuration checkbox.

---

## **2. Using Gateway API**

### **Step 1: Define the Gateway**

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: my-gateway
spec:
  gatewayClassName: gke-l7-global-external-managed # Standard for 2026
  listeners:
    - name: http
      protocol: HTTP
      port: 80
    - name: https
      protocol: HTTPS
      port: 443
      tls:
        mode: Terminate
        certificateRefs:
          - name: my-ssl-cert # Reference to Secret or Certificate Manager
```

---

### **Step 2: Create the HTTP-to-HTTPS Redirect Route**

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: http-to-https-redirect
spec:
  parentRefs:
    - name: my-gateway
      sectionName: http # Targets HTTP listener
  rules:
    - filters:
        - type: RequestRedirect
          requestRedirect:
            scheme: https
            statusCode: 301
```

---

### **Step 3: Create the Main Application Route**

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: app-route
spec:
  parentRefs:
    - name: my-gateway
      sectionName: https # Targets HTTPS listener
  rules:
    - backendRefs:
        - name: my-service
          port: 80
```

**Key Differences from Ingress:**

* **Decoupled Logic:** Listeners and routes are separate, making traffic flow explicit.
* **Standardized:** `RequestRedirect` filter is a standard Kubernetes Gateway API feature.
* **Flexible Status Codes:** Easily choose 301, 302, 307, 308 depending on your needs.

---

✅ This formatting separates **Ingress** vs **Gateway API**, shows **step-by-step YAML**, and highlights **key considerations** for both approaches.

If you want, I can also make a **side-by-side comparison table of Ingress vs Gateway API for HTTP-to-HTTPS redirects**, which is super handy for documentation. Do you want me to do that?
