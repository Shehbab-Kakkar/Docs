`toYaml` is a **Helm template function** used to **convert a map or list into properly formatted YAML**.

It saves you from writing a lot of manual templating.

---

## Why use `toYaml`?

Without `toYaml`:

* You must loop (`range`)
* Handle indentation yourself
* More error-prone 😬

With `toYaml`:

* Helm prints valid YAML automatically
* Cleaner templates
* Safer

---

## Basic Example

### `values.yaml`

```yaml
resources:
  limits:
    cpu: "500m"
    memory: "256Mi"
  requests:
    cpu: "250m"
    memory: "128Mi"
```

---

### Template **without** `toYaml`

```yaml
resources:
  limits:
    cpu: {{ .Values.resources.limits.cpu }}
    memory: {{ .Values.resources.limits.memory }}
  requests:
    cpu: {{ .Values.resources.requests.cpu }}
    memory: {{ .Values.resources.requests.memory }}
```

---

### Template **with** `toYaml` ✅

```yaml
resources:
{{- toYaml .Values.resources | nindent 2 }}
```

---

### Rendered output

```yaml
resources:
  limits:
    cpu: 500m
    memory: 256Mi
  requests:
    cpu: 250m
    memory: 128Mi
```

---

## `toYaml` with `with` (very common)

```yaml
{{- with .Values.resources }}
resources:
{{- toYaml . | nindent 2 }}
{{- end }}
```

✔️ Only renders if `resources` exists
✔️ Clean indentation

---

## Example: labels

### `values.yaml`

```yaml
labels:
  app: my-app
  env: prod
```

### Template

```yaml
metadata:
  labels:
{{- toYaml .Values.labels | nindent 4 }}
```

---

## Example: env variables (list)

### `values.yaml`

```yaml
env:
  - name: LOG_LEVEL
    value: debug
  - name: REGION
    value: us-east-1
```

### Template

```yaml
env:
{{- toYaml .Values.env | nindent 2 }}
```

---

## Important helpers used with `toYaml`

| Function  | Purpose                           |
| --------- | --------------------------------- |
| `toYaml`  | Convert object → YAML             |
| `indent`  | Indent but keep new line          |
| `nindent` | New line + indent (**most used**) |

---

## Common mistake ❌

```yaml
resources: {{ toYaml .Values.resources }}
```

❌ breaks YAML formatting

✅ Correct:

```yaml
resources:
{{- toYaml .Values.resources | nindent 2 }}
```

---

## One-line summary (interview-ready)

> `toYaml` converts Helm values into valid YAML so complex objects can be injected into templates cleanly and safely.

If you want, I can show:

* real **Deployment.yaml** using `toYaml`
* difference between `indent` vs `nindent`
* when **not** to use `toYaml`
