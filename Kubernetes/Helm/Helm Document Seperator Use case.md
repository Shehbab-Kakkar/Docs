Ah — now I see exactly what’s happening. This is a **classic Helm “Namespace object race/behavior” issue**. Let me explain carefully.

---

### 1️⃣ What you did

Your `namespaces.yaml` template is:

```yaml
{{- range .Values.namespaces }}
apiVersion: v1
kind: Namespace
metadata:
  name: {{ . }}
{{- end }}
```

Your `values.yaml`:

```yaml
namespaces:
  - dev
  - staging
  - prod
```

So Helm **generated all three Namespace manifests**, as your dry-run output confirms:

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: dev
apiVersion: v1
kind: Namespace
metadata:
  name: staging
apiVersion: v1
kind: Namespace
metadata:
  name: prod
```

---

### 2️⃣ Why only `prod` shows after install

**Helm installs resources sequentially**, but **Kubernetes cannot create multiple resources in the same manifest if they are not separated properly**.

Look at your dry-run output:

```yaml
---
# Source: namespaces-chart/templates/namespaces.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: dev
apiVersion: v1
kind: Namespace
metadata:
  name: staging
apiVersion: v1
kind: Namespace
metadata:
  name: prod
```

Notice — there are **no `---` document separators** between the namespaces.

Kubernetes interprets **only the last resource (`prod`)** and ignores the rest. That’s why you only see `prod` created.

---

### 3️⃣ Fix: add proper YAML separators

Update `templates/namespaces.yaml` like this:

```yaml
{{- range .Values.namespaces }}
---
apiVersion: v1
kind: Namespace
metadata:
  name: {{ . }}
{{- end }}
```

* The `---` tells Kubernetes each Namespace is a **separate YAML document**
* Helm will now create **all three namespaces properly**

---

### 4️⃣ Test

Dry-run first:

```bash
helm install namespaces-chart ./namespaces-chart --dry-run --debug
```

You should see:

```yaml
---
apiVersion: v1
kind: Namespace
metadata:
  name: dev
---
apiVersion: v1
kind: Namespace
metadata:
  name: staging
---
apiVersion: v1
kind: Namespace
metadata:
  name: prod
```

Then install for real:

```bash
helm upgrade --install namespaces-chart ./namespaces-chart
```

Check:

```bash
kubectl get ns
```

Now you should see:

```
NAME                 STATUS   AGE
default              Active
kube-node-lease      Active
kube-public          Active
kube-system          Active
local-path-storage   Active
dev                  Active
staging              Active
prod                 Active
```

---

✅ **Summary**

* Missing `---` between resources = only last resource applied
* Always separate multiple objects in a single Helm template with `---`
* Now your Helm chart will correctly create **all namespaces**

---

If you want, I can show a **next-level Helm setup** that:

* Deploys the **same chart/service into all three namespaces automatically**
* So you don’t need to run Helm three times for `dev`, `staging`, `prod`

Do you want me to do that?
