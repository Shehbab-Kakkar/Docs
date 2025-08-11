# 🚦 Istio Gateway and VirtualService Configuration

This file explains `istio-gateway.yaml`, which defines how external traffic is routed into the Bookinfo application via Istio.

---

## 🏰 Istio Gateway Resource

- **Name:** `bookinfo-gateway`
- **Selector:** `istio: ingressgateway`
  - Binds this gateway config to pods labeled as the ingress gateway (Envoy).
- **Servers:**  
  - Listens on port 80 (HTTP) for all hosts (`hosts: ["*"]`).

---

## 🧭 Istio VirtualService Resource

- **Name:** `bookinfo`
- **hosts:** `["*"]` (matches all hosts)
- **gateways:** Targets `bookinfo-gateway`
- **http:**  
  - **match:** Routes requests with:
    - `/productpage` (exact)
    - `/static` (prefix)
    - `/login` (exact)
    - `/logout` (exact)
    - `/api/v1/products` (prefix)
  - **route:** Forwards all matching traffic to the `productpage` Service on port 9080.

---

## 🔗 How It Works

1. **Gateway** allows external HTTP traffic on port 80.
2. **VirtualService** matches incoming requests and routes them to the appropriate backend service (here, `productpage`).
3. This enables advanced traffic management, such as:
   - Path-based routing
   - A/B testing and canary releases (with more rules)
   - Integration with Istio security, telemetry, and policy features

---

## 📝 Usage

Apply all resources:
```bash
kubectl apply -f istio-gateway.yaml
```

---

## 🔗 References

- [Istio Gateways](https://istio.io/latest/docs/reference/config/networking/gateway/)
- [Istio VirtualServices](https://istio.io/latest/docs/reference/config/networking/virtual-service/)

---
