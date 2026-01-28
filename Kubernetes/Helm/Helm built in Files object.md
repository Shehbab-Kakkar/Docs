In **Helm**, the built-in **`.Files` object cannot access files in the `templates/` directory**.

### Key rule to remember

👉 **`.Files` only works with non-template files packaged in the chart**, typically things like:

* `files/`
* `config/`
* any other custom directories **outside** `templates/`

### What you *cannot* access with `.Files`

* ❌ **Anything inside `templates/`**
* ❌ Files from **other charts or parent charts** (scope-limited)
* ❌ Files outside the chart package itself

### Why?

Files in `templates/` are rendered as Go templates, not treated as raw files. Helm intentionally excludes them from `.Files` to avoid circular/template-inside-template behavior.

### Example

This **will NOT work**:

```yaml
{{ .Files.Get "templates/config.yaml" }}
```

This **will work**:

```yaml
{{ .Files.Get "files/config.yaml" }}
```

### Common use case

You often use `.Files` to load raw data into:

* ConfigMaps
* Secrets

```yaml
data:
  app.conf: |
{{ .Files.Get "files/app.conf" | indent 4 }}
```

If you want, I can also explain:

* how `.Files` behaves with **subcharts**
* `.Files.Get` vs `.Files.Glob`
* or common Helm exam gotchas 😄
