Perfect! Limiting **open files** in RHEL 9 is a common task for system stability, especially for daemons or database servers. Let’s go **step by step**, including **temporary vs permanent limits**.

---

## 1️⃣ Check current limits

Each process has a limit for open files:

```bash
ulimit -n
```

Example output:

```
1024
```

* This is the **soft limit** for the current shell/session.
* You can see **soft and hard limits**:

```bash
ulimit -Sn   # soft limit
ulimit -Hn   # hard limit
```

---

## 2️⃣ Temporarily change limit (current session only)

```bash
ulimit -n 65535
```

* Now the shell/process can open up to **65535 files**.
* This change is **lost on logout** or reboot.

---

## 3️⃣ Permanently change limits for a user

### Option A: Using `/etc/security/limits.conf`

Edit `/etc/security/limits.conf`:

```bash
sudo vi /etc/security/limits.conf
```

Add lines:

```
username   soft    nofile    65535
username   hard    nofile    65535
```

* `soft` → warning limit (can be increased up to `hard`)
* `hard` → absolute max limit

### Option B: Using `/etc/security/limits.d/`

Create a file, e.g., `/etc/security/limits.d/90-custom.conf`:

```
username   soft    nofile    65535
username   hard    nofile    65535
```

> This is cleaner than editing the main `limits.conf`.

---

## 4️⃣ For systemd-managed services (daemons)

Systemd **ignores PAM limits** by default for services. You must set limits in the **service unit**:

1. Edit the systemd service file:

```bash
sudo systemctl edit --full myservice.service
```

2. Add or modify:

```
[Service]
LimitNOFILE=65535
```

3. Reload systemd and restart the service:

```bash
sudo systemctl daemon-reexec
sudo systemctl restart myservice.service
```

4. Verify limit:

```bash
cat /proc/$(pidof myservice)/limits | grep "Max open files"
```

---

## 5️⃣ Verify limits

* For a shell session:

```bash
ulimit -n
```

* For a running process:

```bash
cat /proc/PID/limits | grep "Max open files"
```

---

### ✅ Summary

| Goal                 | Command / File                                                 |                        |
| -------------------- | -------------------------------------------------------------- | ---------------------- |
| Check current limit  | `ulimit -n`                                                    |                        |
| Temporary limit      | `ulimit -n 65535`                                              |                        |
| Permanent user limit | `/etc/security/limits.conf` or `/etc/security/limits.d/*.conf` |                        |
| Service limit        | `LimitNOFILE=65535` in systemd unit                            |                        |
| Verify for process   | `cat /proc/PID/limits                                          | grep "Max open files"` |

---

💡 **Tip:**

* Always set **hard limit ≥ soft limit**.
* For high-load services like **nginx, database, or Redis**, systemd unit file limits are **mandatory**, because PAM limits don’t apply to background daemons.

---

If you want, I can make a **diagram showing the hierarchy of open file limits** in RHEL 9: **shell → user → PAM → systemd → process**, so you can see which limit takes precedence.

Do you want me to make that diagram?
