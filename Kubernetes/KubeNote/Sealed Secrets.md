Below is a **complete, clean, production-style guide** for installing the **Bitnami Sealed Secrets Controller**, with **proper file names**, **directory structure**, and **exact commands**, followed by a **DB sealed secret example**.

---

# 1. Prerequisites

* Kubernetes cluster (v1.20+ recommended)
* `kubectl`
* `kubeseal` CLI (same major version as controller)

Install `kubeseal` CLI:

```bash
# Linux
wget https://github.com/bitnami-labs/sealed-secrets/releases/latest/download/kubeseal-linux-amd64
chmod +x kubeseal-linux-amd64
sudo mv kubeseal-linux-amd64 /usr/local/bin/kubeseal

# Verify
kubeseal --version
```

---

# 2. Recommended Directory Structure

```text
k8s/
├── sealed-secrets/
│   ├── 01-namespace.yaml
│   ├── 02-controller.yaml
│   └── 03-rbac.yaml
├── secrets/
│   ├── db-secret.yaml
│   └── db-sealedsecret.yaml
└── app/
    └── deployment.yaml
```

---

# 3. Install Sealed Secrets Controller (Manual YAML)

## File: `k8s/sealed-secrets/01-namespace.yaml`

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: kube-system
```

> (Skip if `kube-system` already exists)

---

## File: `k8s/sealed-secrets/02-controller.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sealed-secrets-controller
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      app: sealed-secrets-controller
  template:
    metadata:
      labels:
        app: sealed-secrets-controller
    spec:
      serviceAccountName: sealed-secrets-controller
      containers:
        - name: sealed-secrets-controller
          image: bitnami/sealed-secrets-controller:v0.25.0
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: 8080
          command:
            - controller
```

---

## File: `k8s/sealed-secrets/03-rbac.yaml`

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: sealed-secrets-controller
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: sealed-secrets-controller
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["*"]
  - apiGroups: ["bitnami.com"]
    resources: ["sealedsecrets"]
    verbs: ["*"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: sealed-secrets-controller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: sealed-secrets-controller
subjects:
  - kind: ServiceAccount
    name: sealed-secrets-controller
    namespace: kube-system
```

---

## Apply Controller

```bash
kubectl apply -f k8s/sealed-secrets/
```

Verify:

```bash
kubectl get pods -n kube-system | grep sealed
```

---

# 4. Create a Database Secret (Unsealed)

## File: `k8s/secrets/db-secret.yaml` ❌ **DO NOT COMMIT**

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: db-credentials
  namespace: default
type: Opaque
stringData:
  DB_HOST: postgres.default.svc.cluster.local
  DB_PORT: "5432"
  DB_NAME: appdb
  DB_USER: appuser
  DB_PASSWORD: supersecretpassword
```

---

# 5. Seal the Secret

```bash
kubeseal \
  --controller-name sealed-secrets-controller \
  --controller-namespace kube-system \
  --format yaml \
  < k8s/secrets/db-secret.yaml \
  > k8s/secrets/db-sealedsecret.yaml
```

---

# 6. Sealed Secret (Git-Safe)

## File: `k8s/secrets/db-sealedsecret.yaml` ✅ **SAFE TO COMMIT**

```yaml
apiVersion: bitnami.com/v1alpha1
kind: SealedSecret
metadata:
  name: db-credentials
  namespace: default
spec:
  encryptedData:
    DB_HOST: AgBy8Z9...
    DB_PORT: AjKsd93...
    DB_NAME: Bk29sl...
    DB_USER: Cx92ls...
    DB_PASSWORD: Zx82ls...
  template:
    metadata:
      name: db-credentials
      namespace: default
    type: Opaque
```

---

# 7. Apply the Sealed Secret

```bash
kubectl apply -f k8s/secrets/db-sealedsecret.yaml
```

Check:

```bash
kubectl get secret db-credentials
```

---

# 8. Use in Application

## File: `k8s/app/deployment.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app
spec:
  replicas: 1
  template:
    spec:
      containers:
        - name: app
          image: my-app:latest
          env:
            - name: DB_HOST
              valueFrom:
                secretKeyRef:
                  name: db-credentials
                  key: DB_HOST
            - name: DB_USER
              valueFrom:
                secretKeyRef:
                  name: db-credentials
                  key: DB_USER
            - name: DB_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: db-credentials
                  key: DB_PASSWORD

```

---

# 9. Key Best Practices

✅ Commit **only**:

* `*-sealedsecret.yaml`
* Controller manifests

❌ Never commit:

* Plain `Secret` YAML
* Private controller keys

🔒 Seal scope options:

```bash
--scope strict        # namespace + name
--scope namespace-wide
--scope cluster-wide
```

---

