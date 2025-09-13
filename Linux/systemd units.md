To check all **running systemd units** on a **Rocky Linux 9** server, you can use the `systemctl` command, which is the standard tool for managing `systemd` services.

---

## ✅ Command to List All Running Systemd Units

```bash
systemctl list-units --type=service --state=running
```

### 🔍 Explanation:

* `--type=service`: Limits output to only `*.service` units (systemd services).
* `--state=running`: Shows only services that are **actively running**.

---

## 📋 Example Output

```bash
UNIT                           LOAD   ACTIVE SUB     DESCRIPTION
sshd.service                   loaded active running OpenSSH server daemon
firewalld.service              loaded active running firewalld - dynamic firewall daemon
chronyd.service                loaded active running NTP client/server
...
```

---

## 🧠 Optional Variations

### ✅ List **all services** (not just running):

```bash
systemctl list-units --type=service
```

### ✅ List all **installed** unit files (enabled, disabled, etc.):

```bash
systemctl list-unit-files --type=service
```

### ✅ List all **failed** services:

```bash
systemctl --failed
```

---

## 📌 Summary of Useful Commands

| Task                                   | Command                                               |
| -------------------------------------- | ----------------------------------------------------- |
| Running services only                  | `systemctl list-units --type=service --state=running` |
| All active services (running & exited) | `systemctl list-units --type=service`                 |
| All installed unit files               | `systemctl list-unit-files --type=service`            |
| Failed services                        | `systemctl --failed`                                  |

---

Let me know if you want to filter by user/system scope, or export the output.
