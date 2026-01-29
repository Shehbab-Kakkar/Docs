In Helm, `--atomic` is basically **“all-or-nothing” installs/upgrades**. 💣➡️✅

### What `--atomic` does

When you run **install** or **upgrade** with `--atomic`:

* If **anything fails**, Helm **automatically rolls back**
* The release is left in the **previous working state**
* You don’t get a half-broken deployment

### Typical usage

```bash
helm install myapp ./chart --atomic
```

```bash
helm upgrade myapp ./chart --atomic
```

### What happens under the hood

* Helm waits for resources to become **Ready**
* If readiness fails or a hook fails:

  * **Install** → Helm **uninstalls** the release
  * **Upgrade** → Helm **rolls back** to the last successful revision

### Important details

* `--atomic` **implies `--wait`**

  * You don’t need to specify `--wait` separately
* Strongly recommended to also set a timeout:

```bash
helm upgrade myapp ./chart --atomic --timeout 5m
```

### When to use it

✅ CI/CD pipelines
✅ Production deployments
✅ Any time you want **safety over speed**

### When *not* to use it

⚠️ Debugging charts (rollback hides the broken state)
⚠️ Long-running jobs without proper readiness probes

### Common gotcha

If your pods **never become Ready** (bad probes, missing deps), Helm will:

* Wait until timeout
* Then rollback
* Making it look like “nothing was deployed”

---

If you want, I can also explain:

* How `--atomic` interacts with hooks
* Differences vs `--wait`
* Best practices for probes to avoid false rollbacks
