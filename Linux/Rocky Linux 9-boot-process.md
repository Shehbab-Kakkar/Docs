When you **power on a Rocky Linux 9** (or any modern Linux using `systemd`), the system goes through a well-defined **boot and initialization sequence** to bring up the OS and services.

Here's a clear, step-by-step explanation of **what happens from power-on to a running Rocky 9 system**, specifically with `systemd` involved:

---

## 🧠 High-Level Overview

1. **Firmware (BIOS/UEFI)**
2. **Bootloader (GRUB2)**
3. **Kernel Initialization**
4. **`systemd` takes over (PID 1)**
5. **Target Units & Service Startup**
6. **User Login / Graphical Interface (optional)**

---

## ⚙️ Detailed Boot Steps (Rocky Linux 9 with systemd)

---

### 1. 🔌 **Firmware Initialization (BIOS/UEFI)**

* System is powered on.
* **BIOS or UEFI** performs POST (Power-On Self Test).
* Finds the bootable device (HDD, SSD, etc.).
* Hands over control to the **bootloader** (GRUB).

---

### 2. 🥾 **Bootloader (GRUB2)**

* **GRUB2** loads the **Linux kernel** and **initramfs** (initial RAM filesystem).
* Shows the boot menu (if enabled).
* Passes control to the **Linux kernel**, along with kernel parameters.

---

### 3. 🧬 **Linux Kernel Initialization**

* The **kernel**:

  * Initializes hardware.
  * Mounts the root filesystem (from `initramfs` temporarily, then real root).
  * Starts PID 1: the **init system**, which is `systemd`.

---

### 4. 🧠 **`systemd` Initialization (PID 1)**

* `systemd` is now the **first userspace process** (`PID 1`).
* Reads configuration from:

  * `/etc/systemd/system/`
  * `/lib/systemd/system/`
  * `/etc/systemd/system/default.target` → defines default boot target.

#### Key systemd targets:

| Target Name         | Purpose                                           |
| ------------------- | ------------------------------------------------- |
| `basic.target`      | Prepares basic services (e.g., devices, sockets). |
| `multi-user.target` | Non-graphical multi-user mode (like runlevel 3).  |
| `graphical.target`  | GUI environment (like runlevel 5).                |

---

### 5. 🔄 **Starting Services via Units**

* `systemd` starts all required **unit files**, including:

  * `.service` (daemons)
  * `.mount` (file systems)
  * `.socket` (sockets)
  * `.target` (group of services)

* Units are started **in parallel**, with dependencies resolved using:

  * `After=`
  * `Requires=`
  * `Wants=`

* For example:

  * `network.service`
  * `sshd.service`
  * `firewalld.service`

✅ You can view what was started with:

```bash
systemctl list-units --type=service
```

---

### 6. 👤 **User Login or Graphical Desktop**

* After reaching `multi-user.target` (or `graphical.target`), the system is ready.
* For CLI systems: getty (login prompt) is available.
* For GUI systems: Display manager (like GDM) starts.

---

## 📌 Visual Summary

```plaintext
[Power On]
   ↓
[BIOS/UEFI]
   ↓
[GRUB2 Bootloader]
   ↓
[Linux Kernel]
   ↓
[systemd (PID 1)]
   ↓
[Targets & Units]
   ↓
[Login Prompt or GUI]
```

---

## 🧪 Check Your Boot Process (Optional Commands)

* View boot logs:

```bash
journalctl -b
```

* See startup time per unit:

```bash
systemd-analyze blame
```

* Check boot performance:

```bash
systemd-analyze
```

---

Would you like a graphical diagram or command-line demo to inspect your system's boot on Rocky 9?
