# 📚 Bookinfo Application Kubernetes Resources

This file describes the **Bookinfo microservices application** components as defined in `book-app.yaml`. This application is often used as a sample to demonstrate Istio service mesh features on Kubernetes (such as on GKE).

---

## 🏗️ Microservices Architecture

### Services
- **details**: Provides book details.
- **ratings**: Provides book ratings.
- **reviews**: Three versions (v1, v2, v3) offering book reviews, some of which call the ratings service.
- **productpage**: The main web UI, calls details and reviews.

### For Each Service:
- **Kubernetes Service**: Exposes the app internally in the cluster (ClusterIP).
- **Deployment**: Manages one or more pods for each microservice.
- **ServiceAccount**: Used for pod identity, mTLS, and Istio features.

All Deployments have **Istio sidecar injection enabled** via annotations, allowing traffic between services to flow through Istio’s Envoy proxies for observability, security, and policy enforcement.

---

## 📦 Example: `details` Service

- **Service**: Exposes port 9080, selector `app: details`.
- **Deployment**: One replica, image `istio/examples-bookinfo-details-v1:1.18.0`, annotated for Istio proxy resources, uses `bookinfo-details` ServiceAccount.

---

## 📦 Example: `reviews` Service

- **Service**: Exposes port 9080, selector `app: reviews`.
- **Deployments**: Three versions (`v1`, `v2`, `v3`) each as a separate deployment, annotated for Istio and using `bookinfo-reviews` ServiceAccount.

---

## 🛠️ Istio Integration

- All pods are annotated for Istio sidecar proxy injection.
- Enables advanced service mesh features: traffic routing, telemetry, mTLS, access policies, and more.

---

## 📝 Usage

Apply all resources:
```bash
kubectl apply -f book-app.yaml
```
This will create the Bookinfo services, deployments, and service accounts in your Kubernetes cluster.

---

## 🔗 References

- [Istio Bookinfo Example](https://istio.io/latest/docs/examples/bookinfo/)
- [Kubernetes Services](https://kubernetes.io/docs/concepts/services-networking/service/)
- [Kubernetes Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

---
