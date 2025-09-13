To route **90% of traffic to one version** of an app and **10% to another** using **Istio**, you can use a **VirtualService** to define traffic splitting rules.

Here's a quick guide with an example:

---

### ✅ **Prerequisites**

1. You have Istio installed on your Kubernetes cluster.
2. You have two versions of your app deployed (e.g., `v1` and `v2`) under the same service.

---

### 🧱 **Example Setup**

Assume you have a service named `my-app` with two versions:

* `my-app-v1` (v1)
* `my-app-v2` (v2)

They are both behind a Kubernetes service called `my-app`.

---

### 📄 **VirtualService Configuration**

```yaml
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: my-app
spec:
  hosts:
    - my-app
  http:
    - route:
        - destination:
            host: my-app
            subset: v1
          weight: 90
        - destination:
            host: my-app
            subset: v2
          weight: 10
```

---

### 🔧 **DestinationRule Configuration**

To make the above VirtualService work, define subsets in a **DestinationRule**:

```yaml
apiVersion: networking.istio.io/v1beta1
kind: DestinationRule
metadata:
  name: my-app
spec:
  host: my-app
  subsets:
    - name: v1
      labels:
        version: v1
    - name: v2
      labels:
        version: v2
```

Make sure your pods are labeled correctly. For example:

```yaml
labels:
  app: my-app
  version: v1
```

and

```yaml
labels:
  app: my-app
  version: v2
```

---

### 📌 Apply the configuration

```bash
kubectl apply -f destination-rule.yaml
kubectl apply -f virtual-service.yaml
```

---

### ✅ Verify

Send repeated requests to `my-app` and observe that about 90% go to `v1` and 10% to `v2`. You can do this by adding logging or checking response headers, etc.

---


