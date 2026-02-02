Good catch — this is a **very common MySQL-on-K8s gotcha**, not something you did wrong 👍
Let’s fix it cleanly and I’ll explain *why* it happened so it sticks.

---

## Why this error happens

```text
ERROR 1045 (28000): Access denied for user 'appuser'@'localhost'
```

Key points:

1. **Official MySQL image creates users only on FIRST startup**

   * If the PVC already existed, the env vars were ignored
2. `kubectl exec` connects as **`appuser@localhost`**
3. MySQL auto-created `appuser@'%'`, **not** `appuser@localhost`
4. Result → auth fails

This is expected behavior, not a bug.

---

## ✅ Fastest Fix (recommended)

### Step 1: Login as root (this always works)

```bash
kubectl exec -it mysql-0 -- mysql -uroot -prootpass
```

If this fails, tell me — but it should work.

---

### Step 2: Fix user permissions properly

Run **exactly this** inside MySQL:

```sql
USE appdb;

CREATE USER IF NOT EXISTS 'appuser'@'%' IDENTIFIED BY 'apppass';
GRANT ALL PRIVILEGES ON appdb.* TO 'appuser'@'%';

CREATE USER IF NOT EXISTS 'appuser'@'localhost' IDENTIFIED BY 'apppass';
GRANT ALL PRIVILEGES ON appdb.* TO 'appuser'@'localhost';

FLUSH PRIVILEGES;
```

Now create the table:

```sql
CREATE TABLE IF NOT EXISTS counter (
  id INT PRIMARY KEY,
  visits INT
);

INSERT INTO counter (id, visits)
VALUES (1, 0)
ON DUPLICATE KEY UPDATE id=id;
```

Exit:

```sql
exit;
```

---

### Step 3: Verify appuser works

```bash
kubectl exec -it mysql-0 -- mysql -uappuser -papppass appdb
```

✅ This should now log you in successfully.

---

## 🔧 Optional (but better) fix for Go app DSN

Right now your Go app uses:

```go
tcp(mysql-0.mysql:3306)
```

That **hardcodes one pod**. Better is the service:

```go
dsn := "appuser:apppass@tcp(mysql:3306)/appdb"
```

This lets Kubernetes DNS handle routing.

---

## 🧠 Important MySQL-on-K8s lessons (interview gold)

* MySQL init env vars run **only once**
* PVC reuse = no re-init
* MySQL users are **host-specific**
* `localhost` ≠ `%`
* Always grant both for dev setups

---

## If you want next 🔥

I can:

* Auto-init schema with `initContainers`
* Convert this to **Helm**
* Add **read replicas**
* Replace MySQL with **Vitess / Galera**
* Add **connection pooling + retries** in Go

Just say the word 😄
