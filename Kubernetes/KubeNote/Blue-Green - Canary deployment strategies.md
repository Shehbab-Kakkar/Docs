Kubernetes (K8s) examples for **Blue-Green** and **Canary** deployment strategies:

---

## 1. Blue-Green Deployment

### **Explanation**
- **Blue-Green deployment** involves running two separate environments: "Blue" (current/production) and "Green" (new version).
- At any time, one environment serves all user traffic.
- When you’re ready for an upgrade, you deploy the new version (Green) and then switch all traffic from Blue to Green—either instantly or gradually.
- Rollback is simple: switch back to Blue if Green has issues.

### **Kubernetes Example**

Imagine you have:
- **Blue** deployment: `my-app-blue` (v1)
- **Green** deployment: `my-app-green` (v2)
- **Service** for routing: `my-app-service`

#### **Deployments YAML**

```yaml name=blue-green-deployments.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app-blue
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-app
      version: blue
  template:
    metadata:
      labels:
        app: my-app
        version: blue
    spec:
      containers:
      - name: my-app
        image: my-app:v1

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app-green
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-app
      version: green
  template:
    metadata:
      labels:
        app: my-app
        version: green
    spec:
      containers:
      - name: my-app
        image: my-app:v2
```

#### **Service YAML**

```yaml name=my-app-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: my-app-service
spec:
  selector:
    app: my-app
    version: blue   # Initially points to Blue
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
```

#### **Switching Traffic**
- To shift all users to Green, you change the Service selector to `version: green`.

```yaml
spec:
  selector:
    app: my-app
    version: green
```

---

## 2. Canary Deployment

### **Explanation**
- **Canary deployment** releases the new version to a small percentage of users before exposing it to everyone.
- Traffic is split (e.g., 90% to old version, 10% to new).
- If no issues occur, increase new version traffic gradually.
- Makes impact of bugs much smaller and allows for controlled release.

### **Kubernetes Example**

Imagine:
- **Old deployment**: `my-app-v1` (v1)
- **Canary deployment**: `my-app-v2` (v2)
- **Service**: splits traffic between both.

#### **Deployments YAML**

```yaml name=canary-deployments.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app-v1
spec:
  replicas: 9
  selector:
    matchLabels:
      app: my-app
      version: v1
  template:
    metadata:
      labels:
        app: my-app
        version: v1
    spec:
      containers:
      - name: my-app
        image: my-app:v1

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app-v2
spec:
  replicas: 1
  selector:
    matchLabels:
      app: my-app
      version: v2
  template:
    metadata:
      labels:
        app: my-app
        version: v2
    spec:
      containers:
      - name: my-app
        image: my-app:v2
```

#### **Service YAML**

```yaml name=my-app-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: my-app-service
spec:
  selector:
    app: my-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
```

- The Service sends traffic to both versions. With 9 pods on v1 and 1 pod on v2, **~10% of requests** go to the canary (`v2`).

#### **Increasing Canary Traffic**
- Gradually increase replicas on `my-app-v2` and decrease on `my-app-v1` until all traffic is sent to v2.

---

## **Summary Table**

| Method      | Mechanism                             | Example in K8s                             | Pros/Cons                        |
|-------------|--------------------------------------|--------------------------------------------|----------------------------------|
| Blue-Green  | Two environments, instant switch     | Change Service selector                    | Fast rollback, risk if Green buggy|
| Canary      | Gradual traffic shift, small batch   | Adjust replica counts to split traffic     | Controlled, low blast radius      |

---

Let me know if you want real-world Helm, Argo Rollouts, or Istio examples for advanced use-cases!
