Perfect — you already have a Helm release called `mytest` deployed in the `default` namespace from your `mytest` chart. I’ll break down **install, uninstall, dry-run, and upgrade commands** tailored exactly to your situation.

---

## 1️⃣ **Install the chart**

Since the chart directory is `helm-charts/mytest`:

```bash
helm install mytest ./mytest
```

* `mytest` → release name
* `./mytest` → chart path

> You **already installed it**, so running this again without `--replace` will fail. You only use this the first time.

---

## 2️⃣ **Uninstall the chart**

```bash
helm uninstall mytest
```

* This deletes the release `mytest` from the cluster
* Pods, services, and other resources created by the chart are removed

> After this, `helm list` will no longer show `mytest`.

---

## 3️⃣ **Dry-run / simulate installation**

```bash
helm install mytest ./mytest --dry-run --debug
```

* `--dry-run` → doesn’t create anything, just shows what would happen
* `--debug` → prints detailed manifest output
* Useful to **preview changes before actually deploying**

> You can also do a dry-run for upgrades.

---

## 4️⃣ **Upgrade an existing release**

Suppose you modified `values.yaml` (e.g., changed `service.type` or `nodePort`), you upgrade like this:

```bash
helm upgrade mytest ./mytest
```

* Upgrades the deployed release with **new chart/config**
* If you have changes in `values.yaml`, make sure to include them:

```bash
helm upgrade mytest ./mytest -f values.yaml
```

* Dry-run upgrade:

```bash
helm upgrade mytest ./mytest -f values.yaml --dry-run --debug
```

> This shows what changes will be applied without actually touching the cluster.

---

## 5️⃣ **Quick summary with your current release**

| Action          | Command                                                         |
| --------------- | --------------------------------------------------------------- |
| Install         | `helm install mytest ./mytest`                                  |
| Uninstall       | `helm uninstall mytest`                                         |
| Dry-run install | `helm install mytest ./mytest --dry-run --debug`                |
| Upgrade         | `helm upgrade mytest ./mytest -f values.yaml`                   |
| Dry-run upgrade | `helm upgrade mytest ./mytest -f values.yaml --dry-run --debug` |
| Check status    | `helm list` / `helm status mytest`                              |

---

Port Forword on the VM.

kubectl port-forward  --address 192.168.64.11 svc/mytest-service 30080:80
