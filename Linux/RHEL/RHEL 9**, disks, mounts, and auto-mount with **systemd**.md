Got it! On **RHEL 9**, disks, mounts, and auto-mount points are now tightly integrated with **systemd**, so you can use **systemctl** and other systemd tools to check them. Let me break it down in detail.

---

## 1️⃣ List all mounts

Mount points in systemd are represented as **`.mount` units**.

### Command:

```bash
systemctl list-units --type=mount
```

### Example output:

```
UNIT                  LOAD   ACTIVE SUB     DESCRIPTION
proc.mount            loaded active mounted /proc
sysfs.mount           loaded active mounted /sys
dev-mapper-root.mount loaded active mounted /
home.mount            loaded active mounted /home
run-user-1000.mount   loaded active mounted /run/user/1000
```

* `UNIT` → systemd mount unit
* `ACTIVE` → currently mounted
* `SUB` → actual state (`mounted` if in use)
* `DESCRIPTION` → mount path

---

## 2️⃣ Check **all mount points, including disabled ones**

```bash
systemctl list-unit-files --type=mount
```

Example output:

```
UNIT FILE                  STATE
home.mount                 enabled
mnt-backup.mount           disabled
var-lib-docker.mount       static
```

* `enabled` → auto-mounted at boot
* `disabled` → not auto-mounted
* `static` → required by another unit (cannot be enabled directly)

---

## 3️⃣ Check a **specific mount**

```bash
systemctl status home.mount
```

Example output:

```
● home.mount - /home
     Loaded: loaded (/etc/systemd/system/home.mount; enabled; vendor preset: enabled)
     Active: active (mounted) since Sat 2026-01-24 10:00:00
```

* Shows whether the mount is **active**, **enabled for boot**, and the **unit file path**

---

## 4️⃣ Check **auto-mount points (`.automount`)**

Systemd supports **automatic mounts** via `.automount` units:

### List automount units:

```bash
systemctl list-units --type=automount
```

Example output:

```
UNIT                  LOAD   ACTIVE SUB     DESCRIPTION
home.automount        loaded active waiting Automount for /home
mnt-backup.automount  loaded active waiting Automount for /mnt/backup
```

* `SUB = waiting` → triggers mount on access
* `.automount` unit points to a **corresponding `.mount` unit**

---

### Check a specific automount

```bash
systemctl status home.automount
```

Shows the **mount path**, **unit file**, and **current state**.

---

## 5️⃣ Check **dependencies of a mount**

To see what services depend on a mount:

```bash
systemctl list-dependencies home.mount
```

* Useful if mount fails and some services cannot start.

---

## 6️⃣ Check disks themselves

`systemctl` doesn’t list raw disks (use `lsblk` or `blkid` for that), but you **can check mount units corresponding to disks**:

```bash
systemctl list-units --type=mount | grep /mnt
```

* If `/mnt/data` is a separate partition, systemd will have a `mnt-data.mount` unit.
* You can then check if it’s **enabled at boot** or **auto-mounted**.

---

### ✅ Summary Commands

| Task                    | Command                                     |
| ----------------------- | ------------------------------------------- |
| List active mounts      | `systemctl list-units --type=mount`         |
| List all mount units    | `systemctl list-unit-files --type=mount`    |
| Check specific mount    | `systemctl status <mount>.mount`            |
| List active automounts  | `systemctl list-units --type=automount`     |
| Check automount details | `systemctl status <mount>.automount`        |
| Check dependencies      | `systemctl list-dependencies <mount>.mount` |

---

💡 **Tip:**

* If a mount is **enabled**, it will auto-mount at boot.
* If it’s **disabled**, it will not auto-mount unless triggered manually or via `.automount`.
* `.automount` units can save boot time by **mounting only when accessed**.

---

If you want, I can make a **diagram showing disks → mount units → automount units → services** in RHEL 9.
It will visually show how **systemd handles disks and auto-mounts**.

Do you want me to make that diagram?
