# 🌐 Istio Ingress Gateway Kubernetes Resources

This file explains the Istio Ingress Gateway setup as defined in `ingress-gateway.yaml`. The Ingress Gateway is the entry point for all external traffic into the Istio service mesh on your GKE cluster.

---

## 🚪 Gateway Service

- **Service Name:** `istio-ingressgateway`
- **Type:** `LoadBalancer`
  - Exposes the gateway with a public IP address in GKE.
- **Ports:** 80 (HTTP), 443 (HTTPS)
- **Selector:** `istio: ingressgateway`
  - Routes traffic to pods running the Istio ingress gateway (Envoy proxy).

---

## 🛠️ Gateway Deployment

- **Deployment Name:** `istio-ingressgateway`
- **Pod Labels:** `istio: ingressgateway`
- **Container:** `istio-proxy`
  - Uses the `gateway` injection template.
  - Runs as non-root (user/group 1337) with minimal privileges.
  - Security context set to allow binding privileged ports (80/443).

---

## 🔒 RBAC for TLS Secret Access

- **Role:** Grants read-only access (`get`, `watch`, `list`) to Kubernetes `secrets` in the namespace.
- **RoleBinding:** Binds the Role to the `default` ServiceAccount (can be changed for stricter security).
  - Required for the gateway to dynamically read TLS secrets via Istio’s Secret Discovery Service (SDS).

---

## 📝 Usage

Apply all resources:
```bash
kubectl apply -f ingress-gateway.yaml
```
This will provision the gateway Service, Deployment, and RBAC rules.

---

## 🔗 References

- [Istio Ingress Gateway](https://istio.io/latest/docs/tasks/traffic-management/ingress/ingress-control/)
- [Kubernetes LoadBalancer Service](https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer)
- [Kubernetes RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)

---
