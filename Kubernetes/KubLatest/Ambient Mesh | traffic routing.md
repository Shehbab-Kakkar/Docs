To implement traffic routing with a 10%/90% split between pods (typically for canary deployments or A/B testing) in a Kubernetes cluster using a service mesh like Istio (including Ambient Mesh), you will:

- Create two Kubernetes Services and Deployments (e.g., v1 and v2).
- Use Istio VirtualService to split HTTP (or gRPC) traffic between these versions.

Below is a step-by-step guide with YAML examples:

---

## 1. Example Deployments and Services

Suppose you have two versions of your app: `v1` and `v2`.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-v1
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
      version: v1
  template:
    metadata:
      labels:
        app: myapp
        version: v1
    spec:
      containers:
      - name: myapp
        image: myimage:v1
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-v2
spec:
  replicas: 1
  selector:
    matchLabels:
      app: myapp
      version: v2
  template:
    metadata:
      labels:
        app: myapp
        version: v2
    spec:
      containers:
      - name: myapp
        image: myimage:v2
---
apiVersion: v1
kind: Service
metadata:
  name: myapp
spec:
  selector:
    app: myapp
  ports:
  - port: 80
    targetPort: 8080
```

---

## 2. Istio DestinationRule

Define subsets for each version:

```yaml
apiVersion: networking.istio.io/v1beta1
kind: DestinationRule
metadata:
  name: myapp
spec:
  host: myapp
  subsets:
  - name: v1
    labels:
      version: v1
  - name: v2
    labels:
      version: v2
```

---

## 3. Istio VirtualService for Traffic Splitting

Split traffic 90% to v1, 10% to v2:

```yaml
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: myapp
spec:
  hosts:
  - myapp
  http:
  - route:
    - destination:
        host: myapp
        subset: v1
      weight: 90
    - destination:
        host: myapp
        subset: v2
      weight: 10
```

---

## 4. Apply the Manifests

```sh
kubectl apply -f deployment-v1.yaml
kubectl apply -f deployment-v2.yaml
kubectl apply -f service.yaml
kubectl apply -f destinationrule.yaml
kubectl apply -f virtualservice.yaml
```

---

## 5. Notes

- This works for HTTP/gRPC traffic managed by Istio (including Ambient Mesh with Waypoint Proxy for L7 routing).
- For TCP, use tcp routes in VirtualService.
- If using Ambient Mesh, ensure a Waypoint Proxy is deployed for the namespace.

---

**Summary:**  
Use Istio’s VirtualService and DestinationRule to route 90% of traffic to pods with `version: v1` and 10% to `version: v2`. This is the standard method for canary or progressive delivery on Kubernetes with Istio.

Let me know if you need example manifests for a specific app or have questions about Ambient Mesh!
