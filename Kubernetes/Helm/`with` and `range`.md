Perfect question 👍
`with` and `range` are **Helm (Go template) control structures**. They add **logic** to templates.

I’ll explain **when to use each** and then give **clear examples**.

---

## 1️⃣ `with` — *“work inside an object”*

### Why use `with`?

Use `with` when:

* You want to **avoid repeating long paths**
* You want to **check if a value exists** (it runs only if not empty)

### Example without `with`

```yaml
resources:
  limits:
    cpu: {{ .Values.resources.limits.cpu }}
    memory: {{ .Values.resources.limits.memory }}
```

### Same example **with `with`**

```yaml
{{- with .Values.resources }}
resources:
  limits:
    cpu: {{ .limits.cpu }}
    memory: {{ .limits.memory }}
{{- end }}
```

### What happened?

* `with .Values.resources`
  → sets `.` (dot) to `.Values.resources`
* Inside the block, you use **short paths**
* If `.Values.resources` is **empty or missing**, this block is **skipped**

✅ Cleaner
✅ Safer
✅ Less repetition

---

## 2️⃣ `range` — *“loop over a list or map”*

### Why use `range`?

Use `range` when:

* You have a **list**
* You want to create **multiple YAML entries**

---

### Example: multiple container ports

#### `values.yaml`

```yaml
ports:
  - name: http
    port: 80
  - name: https
    port: 443
```

#### Template using `range`

```yaml
ports:
{{- range .Values.ports }}
  - name: {{ .name }}
    containerPort: {{ .port }}
{{- end }}
```

### Rendered output

```yaml
ports:
  - name: http
    containerPort: 80
  - name: https
    containerPort: 443
```

---

## 3️⃣ `range` with key/value (map)

#### `values.yaml`

```yaml
labels:
  app: my-app
  tier: backend
```

#### Template

```yaml
labels:
{{- range $key, $value := .Values.labels }}
  {{ $key }}: {{ $value }}
{{- end }}
```

---

## 4️⃣ `with` + `range` together (very common)

#### `values.yaml`

```yaml
env:
  - name: LOG_LEVEL
    value: debug
  - name: REGION
    value: us-east-1
```

#### Template

```yaml
{{- with .Values.env }}
env:
{{- range . }}
  - name: {{ .name }}
    value: {{ .value }}
{{- end }}
{{- end }}
```

### Why combine them?

* `with` → checks if `env` exists
* `range` → loops through items

---

## Key things to remember 🧠

| Keyword | Purpose                                  |
| ------- | ---------------------------------------- |
| `with`  | Change context (`.`) and check existence |
| `range` | Loop over list or map                    |
| `.`     | Current object                           |
| `$`     | Root context                             |

Example using `$`:

```yaml
{{- range .Values.containers }}
image: {{ $.Values.image.repository }}:{{ $.Values.image.tag }}
{{- end }}
```

---

## Interview-ready summary

> `with` is used to scope into an object and avoid repeating paths, while `range` is used to iterate over lists or maps to generate repeated YAML blocks.


* whitespace control (`{{-` vs `{{`)
* a **full real Deployment** using both
