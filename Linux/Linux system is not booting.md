If your **Linux system is not booting** (like Rocky Linux 9 or similar RHEL-based distros), the fix depends on **where** the boot process fails.

Here's a **step-by-step troubleshooting guide**, from most common to advanced issues:

---

## 🧭 Step 1: Identify Where It Fails

Watch the screen during boot and note the last message or error.

Typical failure stages:

1. **BIOS/UEFI stage** (hardware)
2. **Bootloader stage** (GRUB)
3. **Kernel loading stage**
4. **Filesystem mounting errors**
5. **systemd service failures**
6. **Login or GUI failures**

---

## 🔧 Troubleshooting by Stage

---

### 🟡 1. System Doesn't Even Start (No GRUB Screen)

**Symptoms**: No GRUB menu, black screen, "no bootable device".

#### 🔍 Check:

* Is disk detected in BIOS/UEFI?
* Secure Boot settings?
* Recent disk changes?

#### ✅ Fix:

* Boot from **Rocky 9 Live CD** or rescue USB.
* Reinstall GRUB:

```bash
# Mount the system
mount /dev/sdXn /mnt       # Replace with your root partition
mount --bind /dev /mnt/dev
mount --bind /proc /mnt/proc
mount --bind /sys /mnt/sys
chroot /mnt

# Reinstall GRUB
grub2-install /dev/sdX     # Replace sdX with your disk (not partition)
grub2-mkconfig -o /boot/grub2/grub.cfg
exit
reboot
```

---

### 🟡 2. GRUB Loads, but Kernel Panics or Errors

**Symptoms**:

* `Kernel panic`
* `unable to mount root fs`
* `initramfs` errors

#### 🔍 Causes:

* Corrupt initramfs or kernel
* Wrong UUID in `/etc/fstab`
* Missing drivers

#### ✅ Fix:

* Boot into **rescue mode** from Rocky ISO
* Mount your root filesystem
* Regenerate `initramfs`:

```bash
chroot /mnt
dracut --force
```

Or install a different kernel if needed.

---

### 🟡 3. Boot Stops at Emergency or Dracut Shell

**Symptoms**:

* Dropped into **dracut** or **emergency** shell
* fstab errors
* Filesystem not found

#### 🔍 Check:

* Run `journalctl -xb` for logs
* Use `blkid` and `lsblk` to inspect devices
* Compare with `/etc/fstab` — any wrong UUIDs?

#### ✅ Fix:

1. Edit `/etc/fstab` and correct or comment broken entries:

```bash
nano /etc/fstab
```

2. Exit or `reboot` after fixing.

---

### 🟡 4. Fails After Reaching systemd (Fails to Load Login/GDM)

**Symptoms**:

* Boot completes but stuck on black screen or login fails
* `Failed to start <service>` messages

#### 🔍 Fix:

* Use **Ctrl+Alt+F2** to switch to a text console.
* Log in and check failed services:

```bash
systemctl --failed
journalctl -xe
```

* Disable or fix the broken service:

```bash
systemctl disable bad-service
```

---

### 🟢 5. Use Rescue Mode (from ISO)

Boot into **Rocky Linux ISO → Troubleshooting → Rescue a Rocky system**

Steps:

```bash
1. Choose "Rescue a Rocky system"
2. Let it mount your system under /mnt/sysimage
3. Run: chroot /mnt/sysimage
4. Now repair like you're in normal Linux
```

---

## 🔒 Common Fixes Summary

| Issue                        | Fix                                                                 |
| ---------------------------- | ------------------------------------------------------------------- |
| GRUB not found               | Reinstall GRUB using rescue mode                                    |
| Kernel panic                 | Regenerate initramfs (`dracut --force`)                             |
| Filesystem errors            | Boot rescue → `fsck /dev/sdXn`                                      |
| Wrong fstab UUIDs            | Correct `/etc/fstab` entries                                        |
| Service failures (e.g., GDM) | Use `systemctl`, `journalctl`, disable or fix the broken service    |
| GUI won't load               | Try `startx`, or reinstall display manager (e.g., `gdm`, `lightdm`) |

---

## 🧪 Advanced Tip: Use `systemd.unit=rescue.target`

If the system boots but fails late in the process:

1. At the GRUB screen, press `e` to edit the boot entry.
2. Add this at the end of the kernel line:

```
systemd.unit=rescue.target
```

3. Press `Ctrl+X` to boot into rescue mode and fix the issue.

---

