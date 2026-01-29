Short answer: **`nindent` = `indent` + new line**.
Long answer with examples 👇

---

## `indent`

👉 **Only adds spaces** to the beginning of each line
👉 **Does NOT add a new line**

### Example

```yaml
labels:{{ toYaml .Values.labels | indent 2 }}
```

### `values.yaml`

```yaml
labels:
  app: my-app
  env: prod
```

### Rendered output ❌

```yaml
labels:  app: my-app
  env: prod
```

❌ **Broken YAML** (everything is on the same line)

---

## `nindent`

👉 Adds a **new line + indentation**
👉 This is why it’s used 90% of the time

### Example

```yaml
labels:
{{ toYaml .Values.labels | nindent 2 }}
```

### Rendered output ✅

```yaml
labels:
  app: my-app
  env: prod
```

---

## Side-by-side comparison

| Function  | Adds newline? | Adds spaces? | Typical use |
| --------- | ------------- | ------------ | ----------- |
| `indent`  | ❌ No          | ✅ Yes        | Inline text |
| `nindent` | ✅ Yes         | ✅ Yes        | YAML blocks |

---

## When should you use `indent`?

Rarely — mostly for **inline content** like annotations with text blocks.

Example:

```yaml
description: {{ .Values.description | indent 2 }}
```

---

## When should you use `nindent`? ✅

Almost always when injecting YAML:

* `labels`
* `annotations`
* `resources`
* `env`
* `affinity`
* `tolerations`

Example:

```yaml
resources:
{{- toYaml .Values.resources | nindent 2 }}
```

---

## Easy memory trick 🧠

* **`nindent` = new line + indent**
* If you’re pasting YAML → **use `nindent`**
* If YAML breaks → you probably used `indent` 😄

---

## Interview one-liner

> `indent` only adds spaces, while `nindent` adds a newline plus spaces, which is why `nindent` is preferred for rendering YAML blocks in Helm templates.

If you want, I can show a **real Deployment snippet** where using `indent` breaks prod 😅

Alright, here’s a **realistic Deployment example** showing how using **`indent` breaks YAML** and **`nindent` fixes it**. This is a *classic prod foot-gun* 😄

---

## ❌ Broken Deployment (using `indent`)

### `values.yaml`

```yaml
resources:
  limits:
    cpu: "500m"
    memory: "256Mi"
```

### `deployment.yaml` (WRONG)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: demo
spec:
  template:
    spec:
      containers:
        - name: app
          image: nginx
          resources:{{ toYaml .Values.resources | indent 12 }}
```

### Rendered output ❌

```yaml
resources:            limits:
              cpu: 500m
              memory: 256Mi
```

🔴 **What goes wrong**

* `indent` does **not** add a new line
* YAML becomes invalid
* Kubernetes error like:

```
error converting YAML to JSON
```

---

## ✅ Fixed Deployment (using `nindent`)

### `deployment.yaml` (CORRECT)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: demo
spec:
  template:
    spec:
      containers:
        - name: app
          image: nginx
          resources:
{{ toYaml .Values.resources | nindent 12 }}
```

### Rendered output ✅

```yaml
resources:
  limits:
    cpu: 500m
    memory: 256Mi
```

✔ Valid YAML
✔ Kubernetes accepts it
✔ No 3 a.m. pager alerts

---

## Why `nindent` works

* Adds a **newline**
* Then adds **12 spaces**
* Perfect for nested Kubernetes structures

---

## Rule of thumb 🔥

> If the line ends with `:` → you almost always need `nindent`

---

## Extra tip (very common pattern)

```yaml
{{- with .Values.resources }}
resources:
{{- toYaml . | nindent 12 }}
{{- end }}
```

This:

* avoids rendering empty blocks
* keeps templates clean
* is considered **Helm best practice**

If you want next:

* a **full production-grade Deployment**
* common Helm interview traps
* or debugging with `helm template`

